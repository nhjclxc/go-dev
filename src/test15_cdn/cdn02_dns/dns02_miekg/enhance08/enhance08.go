package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/miekg/dns"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

//   - 在[enhance06.go](enhance01%2Fenhance06.go)的基础上 支持从配置文件或者配置中心或者数据库热加载域名 → IP 【方便动态更新域名记录，依赖本地权威解析和缓存】
// 也就是取出原来的domainMap，改为使用配置文件或者配置中心或者数据库

// 定义域名记录
var domainMap map[string][]dns.RR = map[string][]dns.RR{
	"local.com.": []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{Name: "local.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
			A:   net.ParseIP("127.0.0.1"),
		},
		&dns.AAAA{
			Hdr:  dns.RR_Header{Name: "local.com.", Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 300},
			AAAA: net.ParseIP("::1"), // ::1 表示ipv6的本地回环地址
		},
		&dns.CNAME{
			Hdr:    dns.RR_Header{Name: "www.local.com.", Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 300},
			Target: "local.com.",
		},
	},
	"*.local.com.": []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{Name: "*.local.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
			A:   net.ParseIP("127.222.222.222"),
		},
	},
}

// 定义上游服务器
var upstreamServer []string = []string{
	"8.8.8.8:53",
	"8.8.8.4:53",
}
var upstreamTimeout = 3 * time.Second

type CacheItem struct {
	msg       *dns.Msg
	expiredAt time.Time
	ttl       time.Duration
}
type enhance06Handler struct {
	cache *expirable.LRU[string, *CacheItem]
}

// genCacheKey 生成缓存key
func (h *enhance06Handler) genCacheKey(qname string, qtype uint16) string {
	return qname + "-" + strconv.Itoa(int(qtype))
}

// addCache 加入缓存
func (h *enhance06Handler) addCache(key string, value *CacheItem) bool {
	return h.cache.Add(key, value)
}

// reAddCache 重新加入缓存
func (h *enhance06Handler) reAddCache(key string, value *CacheItem) bool {
	value.expiredAt = time.Now().Add(value.ttl)
	return h.cache.Add(key, value)
}

// getCache 获取缓存
func (h *enhance06Handler) getCache(key string) (*CacheItem, bool) {
	item, ok := h.cache.Get(key)
	if !ok || item == nil {
		return nil, false
	}
	if time.Now().After(item.expiredAt) {
		// 过期了，在缓存中移除
		h.cache.Remove(key)
		return nil, false
	}
	return item, true
}

func (h *enhance06Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {

	// 创建msg
	msg := &dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	// 判断dns查询是否有内容
	if len(r.Question) == 0 {
		w.WriteMsg(msg)
		return
	}

	// 支持处理多个Question
	for _, question := range r.Question {
		qname := question.Name
		qtype := question.Qtype

		// 检查是否在缓存中
		cacheItem, cacheFound := h.getCache(h.genCacheKey(qname, qtype))
		if cacheFound {
			cacheMsg := cacheItem.msg
			cacheMsg.Id = msg.Id
			msg.Answer = append(msg.Answer, cacheMsg.Answer...)

			// 在将该缓存重新设置，刷新ttl
			h.reAddCache(h.genCacheKey(qname, qtype), cacheItem)

			log.Printf("使用缓存 dns question name：%s, qtype: %d \n", qname, qtype)
			continue
		}
		log.Printf("未使用缓存 dns question name：%s, qtype: %d \n", qname, qtype)

		// 检查是不是自定义的域名
		//rrRecords, found := domainMap[qname]
		_, rrRecords, found := matchDomain(qname)
		if found {
			// 自定义域名
			h.customDnsQuery(msg, rrRecords, qname, qtype)

			minTTL := uint32(300) // 默认 300s
			for _, rr := range rrRecords {
				if rr.Header().Ttl < minTTL {
					minTTL = rr.Header().Ttl
				}
			}
			// 加入缓存
			ttl := time.Duration(minTTL) * time.Second
			h.addCache(h.genCacheKey(qname, qtype), &CacheItem{
				msg:       msg,
				expiredAt: time.Now().Add(ttl),
				ttl:       ttl,
			})

		} else {
			// 递归使用上游dns解析
			upstreamMsg := h.upstreamQuery(qname, qtype)
			if upstreamMsg != nil {
				upstreamMsg.Id = msg.Id
				msg.Answer = append(msg.Answer, upstreamMsg.Answer...)

				// 判断最小的minTtl
				minTtl := uint32(300)
				for _, rr := range upstreamMsg.Answer {
					if rr.Header().Ttl < minTtl {
						minTtl = rr.Header().Ttl
					}
				}
				// 加入缓存
				ttl := time.Duration(minTtl) * time.Second
				h.addCache(h.genCacheKey(qname, qtype), &CacheItem{
					msg:       upstreamMsg,
					expiredAt: time.Now().Add(ttl),
					ttl:       ttl,
				})
			}
		}

	}

	w.WriteMsg(msg)
}

func (h *enhance06Handler) customDnsQuery(msg *dns.Msg, records []dns.RR, qname string, qtype uint16) {
	for _, record := range records {
		if qtype == dns.TypeANY || qtype == record.Header().Rrtype {

			// 直接返回 rrRecord，QUESTION SECTION里面会包含通配符
			//msg.Answer = append(msg.Answer, record)

			// 替换rrRecord里面的name未当前请求的name， 返回的QUESTION SECTION即使当前请求的qname
			rrCopy := dns.Copy(record) // 拷贝原始 RR
			rrCopy.Header().Name = qname
			msg.Answer = append(msg.Answer, rrCopy)
		}
	}
}

// matchDomain *匹配域名
// matchedDomain, rrRecords, found
func matchDomain(qname string) (string, []dns.RR, bool) {
	// 1、域名完全匹配
	rrRecords, found := domainMap[qname]
	if found {
		return qname, rrRecords, true
	}

	// 2、*匹配
	for domain, rrs := range domainMap {
		if strings.HasPrefix(domain, "*") {
			// *.local.com. *.example.com. ...
			// 转化为：local.com.  example.com.
			baseDomain := domain[2:]
			if strings.HasSuffix(qname, baseDomain) {
				return domain, rrs, true
			}
		}
	}

	return "", nil, false
}

func (h *enhance06Handler) upstreamQuery(qname string, qtype uint16) *dns.Msg {
	upstreamMsg := &dns.Msg{}
	upstreamMsg.RecursionAvailable = true
	upstreamMsg.SetQuestion(qname, qtype)

	for _, server := range upstreamServer {
		client := dns.Client{}
		exchange, _, err := client.Exchange(upstreamMsg, server)
		if err != nil {
			fmt.Printf("递归上游dns服务器【%s】出错：%s \n", server, err)
			continue
		}
		return exchange
	}

	return nil
}

func main() {

	// 创建处理器
	handler := &enhance06Handler{
		cache: expirable.NewLRU[string, *CacheItem](128, nil, 0),
	}

	// tcp，udp的dns服务
	tcpDnsServer := dns.Server{
		Addr:         ":53",
		Net:          "tcp",
		Handler:      handler,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	udpDnsServer := dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: handler,
	}
	eg := errgroup.Group{}

	eg.Go(func() error {
		log.Println("TCP DNS server listening on :53...")
		return tcpDnsServer.ListenAndServe()
	})
	eg.Go(func() error {
		log.Println("UDP DNS server listening on :53...")
		return udpDnsServer.ListenAndServe()
	})

	if err := eg.Wait(); err != nil {
		log.Fatal("dns服务器启动异常：", err)
	}
}

// dig @127.0.0.1 local.com A a.local.com A bb.local.com A
// dig @"127.0.0.1" '*.local.com' A
