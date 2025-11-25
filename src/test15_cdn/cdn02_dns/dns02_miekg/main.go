package main

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
)

// 使用"github.com/miekg/dns"包实现dns服务器解析，将local.com解析为127.0.0.1
func main() {

	// 1、创建dns服务
	dnsServer := &dns.Server{
		Addr: ":53",
		Net:  "udp",
	}

	// 2、定义dns处理器
	dnsServer.Handler = &CustomDNSHandler{}

	// 3、启动dns服务器
	log.Println("DNS server listening on :53...")
	err := dnsServer.ListenAndServe()
	if err != nil {
		fmt.Println("dns服务启动失败：", err)
		return
	}

}

type CustomDNSHandler struct {
}

// ServeDNS 实现接口已实现DNSHandler
func (h *CustomDNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	// 1、创建消息响应
	msg := &dns.Msg{}
	msg.SetReply(r)

	// 没有dns查询信息直接返回
	if len(r.Question) == 0 {
		w.WriteMsg(msg)
		return
	}

	// 获取请求域名
	// r.Question表示 DNS 查询报文中的 Question Section（问题部分），它包含了客户端发来的 查询域名 + 查询类型（A、AAAA、SRV 等）
	qname := r.Question[0].Name
	qtype := r.Question[0].Qtype
	log.Printf("Query: %s %d\n", qname, qtype)

	// DNS 域名后一定要带点【完全限定域名必须以点.结尾，表示这是 完整域名，没有默认附加任何父域。】
	if "local.com." == qname && qtype == dns.TypeA {
		// 是自定义域名，且是a记录请求，则进入自定义处理程序
		fmt.Println("进入 local.com")

		// 构造自定义域名响应
		rr := &dns.A{
			Hdr: dns.RR_Header{
				Name:   qname,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    300,
			},
			A: net.ParseIP("127.0.0.1"),
		}
		// 将目标ip写到响应里面
		msg.Answer = append(msg.Answer, rr)

		// 表示这是权威 DNS：
		msg.Authoritative = true
		log.Printf("Responding to %s with %v\n", qname, msg.Answer)

		// 将解析结果响应给客户端
		w.WriteMsg(msg)
	}

}

// dig @127.0.0.1 local.com
// dig @127.0.0.1 example.com
// dig @127.0.0.1 google.com

/*
dig @127.0.0.1 local.com A
// dig：（Domain Information Groper）
// @127.0.0.1：明确指定当前的dig命令使用自己本地的 DNS 服务器，不加 @127.0.0.1 → 系统默认 DNS → 返回公网 IP
// local.com：要查询的域名（Query Name），dig 会把这个域名封装到 DNS Query 的 Question Section 中发送给服务器
// A：查询类型（Query Type / QTYPE），A 表示 IPv4 地址记录


lxc20250729@lxc20250729deMacBook-Pro ~ % dig @127.0.0.1 local.com A

; <<>> DiG 9.10.6 <<>> @127.0.0.1 local.com A
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 19052
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;local.com.			IN	A

;; ANSWER SECTION:
local.com.		300	IN	A	127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 11:20:06 CST 2025
;; MSG SIZE  rcvd: 52


解析以下这部分权威dns服务器返回给客户端的内容
```
;; ANSWER SECTION:
local.com.		300	IN	A	127.0.0.1
```
对应 一个 DNS 资源记录 (RR, Resource Record)，一共有 5 个字段：
| 字段                | 示例           | 含义                                                                      |
| ----------------- | ------------ | ----------------------------------------------------------------------- |
| **Name**          | `local.com.` | 域名。这里是被解析的目标域名。末尾的 `.` 表示这是 **完全限定域名 (FQDN)**，不依赖任何搜索域。                 |
| **TTL**           | `300`        | Time To Live，缓存时间（单位：秒）。客户端或中间 DNS 可以缓存这个记录 300 秒。                      |
| **Class**         | `IN`         | DNS 类别，这里是 **IN = Internet**。常见还有 `CH` (Chaos)、`HS` (Hesiod)，几乎都用 `IN`。 |
| **Type**          | `A`          | 记录类型，这里是 **A 记录**，表示 IPv4 地址。其他类型还可能有 AAAA、CNAME、MX 等。                  |
| **RData / Value** | `127.0.0.1`  | 记录的具体值。对于 A 记录，就是对应的 IPv4 地址；对于 AAAA 是 IPv6；CNAME 是指向的域名。               |

这一条记录的意思就是：域名 local.com. 对应 IPv4 地址 127.0.0.1，这个信息可以缓存 300 秒，属于互联网类记录 (IN)，类型是 A。



*/
