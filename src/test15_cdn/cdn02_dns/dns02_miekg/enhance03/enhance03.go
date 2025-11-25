package main

import (
	"github.com/miekg/dns"
	"log"
	"net"
	"time"
)

//- 在[enhance02.go](enhance02%2Fenhance02.go)基础上 支持多 Question 【一个报文可以同时请求多个域名，提高客户端兼容性】

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
	dnsServer.Handler = &enhance03DNSHandler{}

	// 3、启动dns服务器
	log.Printf("Authoritative DNS Server started on 53 port")
	if err := dnsServer.ListenAndServe(); err != nil {
		log.Fatal("DNS Server start error : ", err)
	}
}

// 定义自定义dns处理器，以实现dns的 ServeDNS(w dns.ResponseWriter, r *dns.Msg) 接口
type enhance03DNSHandler struct{}

func (h enhance03DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
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

		// 3、判断是不是自定义的域名
		rrRecords, found := domainMap[qname]
		if found {
			// 自定义的
			h.customDNSQuery(msg, qname, qtype, rrRecords)

		} else {
			// 使用系统的，进行上游递归查询
			// 上游递归查询dns要构造一个新的msg，而本地权威dns就不需要构造一个新的msg，这是为什么？
			//	权威解析：直接在客户端请求 Msg 上填 Answer 即可，不必新建 Msg
			//	递归转发：需要新建 Msg 向上游发送，保证请求干净、控制标志，并在返回时修正 ID
			upstreamResp := h.upstreamDNSQuery(qname, qtype)
			if upstreamResp != nil {
				upstreamResp.Id = msg.Id
				msg.Answer = append(msg.Answer, upstreamResp.Answer...)
			}
		}
	}

	// 响应查询结果
	w.WriteMsg(msg)
	return

}

// customDNSQuery 本地权威dns处理
func (h enhance03DNSHandler) customDNSQuery(msg *dns.Msg, qname string, qtype uint16, rrRecords []dns.RR) {
	for _, rr := range rrRecords {
		if qtype == dns.TypeANY || qtype == rr.Header().Rrtype {
			msg.Answer = append(msg.Answer, rr)
		}
	}
}

// upstreamDNSQuery 递归查询上游服务器
func (h enhance03DNSHandler) upstreamDNSQuery(qname string, qtype uint16) *dns.Msg {
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

// dig @127.0.0.1 local.com A local.com AAAA

/*

lxc20250729@lxc20250729deMacBook-Pro ~ % dig @127.0.0.1 local.com A local.com AAAA

; <<>> DiG 9.10.6 <<>> @127.0.0.1 local.com A local.com AAAA
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 14936
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;local.com.			IN	A

;; ANSWER SECTION:
local.com.		300	IN	A	127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 15:15:19 CST 2025
;; MSG SIZE  rcvd: 52

;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 15392
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;local.com.			IN	AAAA

;; ANSWER SECTION:
local.com.		300	IN	AAAA	::1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 15:15:19 CST 2025
;; MSG SIZE  rcvd: 64

*/
