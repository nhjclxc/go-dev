package main

import (
	"fmt"
	"github.com/miekg/dns"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"strings"
	"time"
)

// 在[main.go](main.go)基础上监听 TCP 和 UDP（兼容大包）【保证服务器能接收请求，UDP 用于常规查询，TCP 用于大包和区分】
// 要想同时监听tcp和udp，因此就必须开两个协程分别监听，因为tcp或udp监听的时候是阻塞监听的

// UDP 默认查询（快），但报文上限 512 字节（没开 EDNS0）
// TCP 大包、DNSSEC、AXFR、客户端强制使用（如 dig +tcp）
//因此一个完整 DNS Server 必须同时监听 UDP 和 TCP。【【【详细看enhance012.go】】】

// 支持 TCP 超时：设置 ReadTimeout / WriteTimeout
// 日志分离：区分 UDP / TCP 请求日志
// AXFR / Zone Transfer：TCP 是必须的

var domainMap = map[string]string{
	"local.com.":   "127.0.0.1",
	"example.com.": "192.168.201.16",
}

type tcpudpDNSHandler struct{}

func (h *tcpudpDNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {

	netType := w.RemoteAddr().Network()
	if strings.HasPrefix(netType, "tcp") {
		fmt.Println("当前是 TCP DNS 请求")
	} else if strings.HasPrefix(netType, "udp") {
		fmt.Println("当前是 UDP DNS 请求")
	} else {
		fmt.Println("未知类型")
	}

	// 1、创建msg返回消息
	msg := &dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	// 2、检查dns查询里面是否有内容
	if len(r.Question) == 0 {
		w.WriteMsg(msg)
		return
	}

	// 3、拿出第一个dns查询内容
	question := r.Question[0]
	qname := question.Name
	qtype := question.Qtype

	// 4、判断qname是不是自定义的域名
	if ip, ok := domainMap[qname]; ok {
		// 是定义的域名，则使用当前的权威域名服务器进行ip地址查询
		if qtype == dns.TypeA {
			rr := &dns.A{
				Hdr: dns.RR_Header{
					Name:   qname,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    300,
				},
				A: net.ParseIP(ip),
			}

			// 追加响应结果
			msg.Answer = append(msg.Answer, rr)
			// 返回查询结果数据
			w.WriteMsg(msg)
			return
		}

	} else {
		// ...使用系统的dns查询
	}

	// 没找到 → NXDOMAIN
	msg.Rcode = dns.RcodeNameError
	w.WriteMsg(msg)
}

func main() {
	// 1、创建处理器
	handler := &tcpudpDNSHandler{}

	// 2、创建tcp监听
	//| 超时类型             | 作用                       |
	//| ---------------- | ------------------------ |
	//| **ReadTimeout**  | 限制“客户端多久不发送数据就断开”，防止慢速攻击 |
	//| **WriteTimeout** | 限制“服务器多久写回应超时”，避免卡死      |

	tcpDnsServer := &dns.Server{
		Addr:    ":53",
		Net:     "tcp",
		Handler: handler,
		// 支持 TCP 超时：设置 ReadTimeout / WriteTimeout
		ReadTimeout:  5 * time.Second, // 读超时
		WriteTimeout: 5 * time.Second, // 写超时
	}
	udpDnsServer := &dns.Server{Addr: ":53", Net: "udp", Handler: handler}

	// 启动两个协程去监听
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
		log.Fatal("dns服务出错：", err)
		return
	}
}

// 测试udp：dig @127.0.0.1 local.com A
// 测试tcp：dig @127.0.0.1 local.com A +tcp

/*
lxc20250729@lxc20250729deMacBook-Pro ~ % dig @127.0.0.1 local.com A

; <<>> DiG 9.10.6 <<>> @127.0.0.1 local.com A
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 1784
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;local.com.			IN	A

;; ANSWER SECTION:
local.com.		300	IN	A	127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 14:55:12 CST 2025
;; MSG SIZE  rcvd: 52

lxc20250729@lxc20250729deMacBook-Pro ~ % dig @127.0.0.1 local.com A +tcp

; <<>> DiG 9.10.6 <<>> @127.0.0.1 local.com A +tcp
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 8535
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;local.com.			IN	A

;; ANSWER SECTION:
local.com.		300	IN	A	127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov 24 14:55:14 CST 2025
;; MSG SIZE  rcvd: 52


*/
