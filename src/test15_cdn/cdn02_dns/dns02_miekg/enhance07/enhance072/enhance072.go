package main

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"time"
)

// 在[test15_cdn/cdn02_dns/dns02_miekg/enhance01.go]的基础上实现【var domainMap = map[string]dns.RR{】的预定义

// DNS 域名后一定要带点【完全限定域名必须以点.结尾，表示这是 完整域名，没有默认附加任何父域。】
var domainMap = map[string]dns.RR{
	"local.com.": &dns.A{
		Hdr: dns.RR_Header{Name: "local.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
		A:   net.ParseIP("127.0.0.1"),
	},
	"example.com.": &dns.A{
		Hdr: dns.RR_Header{Name: "example.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
		A:   net.ParseIP("192.168.201.16"),
	},
}

// 上游 DNS 列表，可增加多个备用 DNS
var upstreamServers = []string{
	"8.8.8.8:53",
	"8.8.4.4:53",
}
var upstreamTimeout = 3 * time.Second

// 在main.go的基础上支持递归查询
func main() {

	// 1、创建dns服务器
	dnsServer := &dns.Server{
		Addr: ":53",
		Net:  "udp",
	}

	// 2、定义自定义dns处理器
	dnsServer.Handler = &enhance01DNSHandler{}

	// 3、启动dns服务器
	log.Printf("Authoritative DNS Server started on 53")
	if err := dnsServer.ListenAndServe(); err != nil {
		log.Fatal("DNS Server start error : ", err)
	}

}

// 定义自定义dns处理器，以实现dns的 ServeDNS(w ResponseWriter, r *Msg) 接口
type enhance01DNSHandler struct{}

// ServeDNS 自定义的dns处理方法
// Params:
//   - dns.ResponseWriter: 用来向客户端写入 DNS 响应，也可以获取客户端信息
//   - *dns.Msg: 表示客户端发送过来的 DNS 请求报文
func (h *enhance01DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	// 1、创建当前这个权威dns服务器的响应消息
	msg := &dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	// 2、判断dns请求里面是否有查询内容，没有查询内容直接返回
	if len(r.Question) == 0 {
		w.WriteMsg(msg)
		return
	}

	// 3、取出第一个r.Question查询，（r.Question支持多个dns查询，但是实际客户端每一次只会发送一个dns查询，所以每次都取出第一个dns查询进行处理即可）
	// 一个 DNS 查询报文（dns.Msg）里面，Question Section 存放客户端请求的 查询条目（Question）。
	// r.Question 是一个 slice：[]dns.Question，每个元素都是一个查询，
	question := r.Question[0]
	qname := question.Name
	qtype := question.Qtype
	fmt.Printf("dns question name：%s, qtype: %d \n", qname, qtype)

	// 4、检测当前查询的域名是不是我们要处理的域名
	rr, ok := domainMap[qname]
	if ok {
		if dns.TypeA == qtype {
			// 5、是我们要处理的域名，则构造一个A记录返回，进行返回指定的ip
			// 6、追加rr响应到当前要返回的消息msg中
			msg.Answer = append(msg.Answer, rr)
			log.Printf("Responding to %s with %v\n", qname, msg.Answer)
		}
	} else {
		// 7、不是当前要处理的域名，直接递归调用上层dns处理器【也就是【支持 递归查询 / 上游转发（像 dnsmasq）】】
		upstreamMsg := h.queryUpstream(qname, qtype)
		if upstreamMsg != nil {
			// ;; Warning: ID mismatch: expected ID 55538, got 33292【这时上游返回的id和客户端请求的id不匹配冲突，只需要将id设置为与客户端请求一至即可】
			// 在返回客户端之前，把上游响应的 ID 改成客户端请求的 ID：
			upstreamMsg.Id = r.Id // 使用客户端的 ID
			w.WriteMsg(upstreamMsg)
			return
		}
	}
	// 8、返回解析结果， 有可能是空
	w.WriteMsg(msg)
}

// queryUpstream 实现【支持 递归查询 / 上游转发（像 dnsmasq）】
func (h *enhance01DNSHandler) queryUpstream(qname string, qtype uint16) *dns.Msg {
	// 创建msg对象
	msg := &dns.Msg{}
	msg.SetQuestion(qname, qtype)
	msg.RecursionDesired = true // 允许上游进行递归查询

	// 遍历上游dns服务器进行查询，只要找到了就立马退出
	for _, server := range upstreamServers {
		client := &dns.Client{Timeout: upstreamTimeout}
		exchangeMsg, _, err := client.Exchange(msg, server)
		if err != nil {
			log.Printf("上游 %s 查询失败: %v\n", server, err)
			continue // 接着查询
		}
		exchangeMsg.Authoritative = false // 上游返回的为非权威的

		// 找到了直接返回
		return exchangeMsg
	}

	return nil
}

/*
dns.ResponseWriter: 用来向客户端写入 DNS 响应，也可以获取客户端信息
| 方法                       | 说明                              |
| ------------------------ | ------------------------------- |
| `WriteMsg(msg *dns.Msg)` | 将一个 `dns.Msg` 发送给客户端，通常在处理完成后调用 |
| `Write([]byte)`          | 直接写原始字节数据给客户端                   |
| `Close()`                | 关闭连接                            |
| `LocalAddr()`            | 返回本地监听地址（服务器端 IP:Port）          |
| `RemoteAddr()`           | 返回客户端地址（客户端 IP:Port）            |
| `TsigStatus()`           | TSIG 校验状态（可选）                   |
| `Hijack()`               | 拿到底层 net.Conn 进行自定义处理           |


*dns.Msg: 表示客户端发送过来的 DNS 请求报文
| 字段                          | 说明                                 |
| --------------------------- | ---------------------------------- |
| `r.Question []dns.Question` | Question Section，客户端查询的域名和类型       |
| `r.Id uint16`               | 请求 ID，用于匹配响应                       |
| `r.RecursionDesired bool`   | RD 标志，客户端请求递归查询                    |
| `r.Opcode int`              | 操作码（QUERY、IQUERY、STATUS）           |
| `r.AuthenticatedData bool`  | AD 标志                              |
| `r.CheckingDisabled bool`   | CD 标志                              |
| `r.Truncated bool`          | TC 标志                              |
| `r.MsgHdr`                  | 其他 Header 信息（Flags、ResponseCode 等） |

*/

// dig @127.0.0.1 local.com A
// dig @127.0.0.1 google.com A

/*
lxc20250729@lxc20250729deMacBook-Pro ~ % dig @127.0.0.1 local.com A

; <<>> DiG 9.10.6 <<>> @127.0.0.1 local.com A
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 7008
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;local.com.			IN	A

;; ANSWER SECTION:
local.com.		300	IN	A	127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 11:55:30 CST 2025
;; MSG SIZE  rcvd: 52



// 没加：queryUpstream
lxc20250729@lxc20250729deMacBook-Pro ~ % dig @127.0.0.1 google.com A

; <<>> DiG 9.10.6 <<>> @127.0.0.1 google.com A
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 5072
;; flags: qr aa rd; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;google.com.			IN	A

;; Query time: 31 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 12:18:32 CST 2025
;; MSG SIZE  rcvd: 28



// 加了：queryUpstream

lxc20250729@lxc20250729deMacBook-Pro ~ % dig @127.0.0.1 google.com A

; <<>> DiG 9.10.6 <<>> @127.0.0.1 google.com A
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 18616
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;google.com.			IN	A

;; ANSWER SECTION:
google.com.		118	IN	A	142.250.71.174

;; Query time: 35 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 12:16:18 CST 2025
;; MSG SIZE  rcvd: 54

lxc20250729@lxc20250729deMacBook-Pro ~ %

*/
