package main

import (
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/miekg/dns"
	"log"
	"net"
	"strconv"
	"time"
)

// enhance05对每一个key都缓存了相同的时间，这里做一个改进，对当前权威域名缓存300s，上游返回的缓存900s

// 自定义域名记录
var domainMap = map[string][]dns.RR{
	"local.com.": []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{Name: "local.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
			A:   net.ParseIP("127.0.0.1"),
		},
		&dns.AAAA{
			Hdr:  dns.RR_Header{Name: "local.com.", Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 300},
			AAAA: net.ParseIP("::1")},
		&dns.CNAME{
			Hdr:    dns.RR_Header{Name: "www.local.com.", Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 300},
			Target: "local.com.",
		},
	},
	"example.com.": []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{Name: "example.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
			A:   net.ParseIP("127.0.0.222"),
		},
	},
}

// 上游dns服务器
var upstreamServer []string = []string{
	"8.8.8.8:53",
	"8.8.8.4:53",
}
var upstreamTimeout time.Duration = 3 * time.Second

func main() {
	//1、定义dns权威服务器
	dnsServer := dns.Server{
		Addr: ":53",
		Net:  "udp",
	}

	// 2、自定义dns处理器
	dnsServer.Handler = &enhance05DNSHandler{
		// 使用 go get github.com/hashicorp/golang-lru/v2来进行dns查询缓存
		// size=128是缓存大小
		// onEvict是内存被淘汰后的回调函数
		// 如果自定义缓存过期，每一个key的过期时间不一样的化，ttl必须设置为0
		cache: expirable.NewLRU[string, *CacheItem](128, nil, 0),
	}

	// 3、启动dns服务器
	log.Printf("Authoritative DNS Server started on 53 port")
	if err := dnsServer.ListenAndServe(); err != nil {
		log.Fatal("DNS Server start error : ", err)
	}
}

// 定义自定义dns处理器，以实现dns的 ServeDNS(w dns.ResponseWriter, r *dns.Msg) 接口
type enhance05DNSHandler struct {
	// key = qname + qtype
	cache *expirable.LRU[string, *CacheItem]
}

type CacheItem struct {
	Msg      *dns.Msg
	ExpireAt time.Time
}

// AddCache 将*dns.Msg进行缓存
func (h enhance05DNSHandler) AddCache(key string, value *CacheItem) bool {
	return h.cache.Add(key, value)
}

// GetCache get缓存时自己判断是否该过期，即每一个key的过期时间不同
func (h enhance05DNSHandler) GetCache(key string) (*CacheItem, bool) {
	item, ok := h.cache.Get(key)
	if !ok || item == nil {
		return nil, false
	}
	// 判断是否过期
	if time.Now().After(item.ExpireAt) {
		h.cache.Remove(key)
		return nil, false
	}
	return item, true
}

func (h enhance05DNSHandler) genCahceKey(qname string, qtype uint16) string {
	return qname + "-" + strconv.Itoa(int(qtype))
}
func (h enhance05DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	// 1、创建msg消息
	msg := &dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	// 2、检查dns查询里面是否有内容
	if len(r.Question) == 0 {
		w.WriteMsg(msg)
		return
	}

	// 支持多个Question的处理
	for _, question := range r.Question {
		// 依次处理每个 Question

		qname := question.Name
		qtype := question.Qtype

		// =========
		//    优先访问缓存的内容
		// =========
		cacheItem, chaheFound := h.GetCache(h.genCahceKey(qname, qtype))
		if chaheFound {
			// 找到记录了，直接用缓存里面的
			cacheMsg := cacheItem.Msg
			cacheMsg.Id = msg.Id
			msg.Answer = append(msg.Answer, cacheMsg.Answer...)

			log.Printf("使用缓存 dns question name：%s, qtype: %d \n", qname, qtype)

			// 继续处理下一个查询
			continue
		}

		log.Printf("未使用缓存 dns question name：%s, qtype: %d \n", qname, qtype)

		// 3、判断是不是自定义的域名
		rrRecords, found := domainMap[qname]
		if found {
			// 自定义的
			h.customDNSQuery(msg, qname, qtype, rrRecords)

			// 查询到的msg加入缓存， msg.Copy()对消息的所有字段做 深拷贝
			//h.AddCache(h.genCahceKey(qname, qtype), &CacheItem{Msg: msg.Copy(), ExpireAt: time.Now().Add(10 * time.Second)})

			// 根据上游返回 RR 动态计算 TTL
			minTTL := uint32(300) // 默认 300s
			for _, rr := range rrRecords {
				if rr.Header().Ttl < minTTL {
					minTTL = rr.Header().Ttl
				}
			}
			h.AddCache(h.genCahceKey(qname, qtype), &CacheItem{Msg: msg.Copy(), ExpireAt: time.Now().Add(time.Duration(minTTL) * time.Second)})

		} else {
			// 使用系统的，进行上游递归查询
			upstreamResp := h.upstreamDNSQuery(qname, qtype)
			if upstreamResp != nil {
				upstreamResp.Id = msg.Id
				msg.Answer = append(msg.Answer, upstreamResp.Answer...)

				// 查询到的msg加入缓存
				//h.AddCache(h.genCahceKey(qname, qtype), &CacheItem{Msg: upstreamResp.Copy(), ExpireAt: time.Now().Add(5 * time.Second)})

				// 根据上游返回 RR 动态计算 TTL
				minTTL := uint32(300) // 默认 300s
				for _, rr := range upstreamResp.Answer {
					if rr.Header().Ttl < minTTL {
						minTTL = rr.Header().Ttl
					}
				}
				h.AddCache(h.genCahceKey(qname, qtype), &CacheItem{Msg: upstreamResp.Copy(), ExpireAt: time.Now().Add(time.Duration(minTTL) * time.Second)})

			}
		}
	}

	// 响应查询结果
	w.WriteMsg(msg)
	return

}

// 本地权威域名和上游递归返回的域名缓存时长选择
// | 类型                              | TTL 建议                                 | 原因                                                  |
//| ------------------------------------ | -------------------------------------- | --------------------------------------------------- |
//| 本地权威域名（local.com / example.com） | 长 TTL，例如 300s~3600s                    | 权威解析固定 IP，不会频繁变动，缓存久一些节省查询                          |
//| 上游递归返回的域名                       | 使用上游返回的 TTL（Answer 中的 `RR_Header.Ttl`） | 上游 DNS 已经给了官方 TTL，遵循标准 DNS 行为；一般比权威 TTL 短，缓存时间按官方给定 |

// customDNSQuery 本地权威dns处理
func (h enhance05DNSHandler) customDNSQuery(msg *dns.Msg, qname string, qtype uint16, rrRecords []dns.RR) {
	for _, rr := range rrRecords {
		if qtype == dns.TypeANY || qtype == rr.Header().Rrtype {
			msg.Answer = append(msg.Answer, rr)
		}
	}
}

// upstreamDNSQuery 递归查询上游服务器
func (h enhance05DNSHandler) upstreamDNSQuery(qname string, qtype uint16) *dns.Msg {
	upStreamMsg := &dns.Msg{}
	upStreamMsg.SetQuestion(qname, qtype)
	upStreamMsg.RecursionAvailable = true // 允许上游递归

	for _, server := range upstreamServer {
		client := dns.Client{Timeout: upstreamTimeout}
		exchangeMsg, _, err := client.Exchange(upStreamMsg, server)
		if err != nil {
			log.Printf("系统dns服务器【%s】查询失败：%s \n", server, err)
			continue // 继续查询下一个服务器
		}
		exchangeMsg.Authoritative = false // 上游返回的为非权威的

		// 查询成功了，直接返回
		return exchangeMsg
	}

	return nil
}

// dig @127.0.0.1 local.com A google.com A

//2025/11/24 16:30:58 未使用缓存 dns question name：local.com., qtype: 1
//2025/11/24 16:30:58 未使用缓存 dns question name：google.com., qtype: 1

//2025/11/24 16:31:01 使用缓存 dns question name：local.com., qtype: 1
//2025/11/24 16:31:01 使用缓存 dns question name：google.com., qtype: 1

//2025/11/24 16:31:05 使用缓存 dns question name：local.com., qtype: 1
//2025/11/24 16:31:05 未使用缓存 dns question name：google.com., qtype: 1

//2025/11/24 16:31:10 未使用缓存 dns question name：local.com., qtype: 1
//2025/11/24 16:31:10 未使用缓存 dns question name：google.com., qtype: 1

//2025/11/24 16:31:14 使用缓存 dns question name：local.com., qtype: 1
//2025/11/24 16:31:14 使用缓存 dns question name：google.com., qtype: 1
