package main

import (
	"log"
	"net"
	"time"

	"github.com/miekg/dns"
)

var zoneRecords = []dns.RR{
	// 必须以 SOA 开头
	&dns.SOA{
		Hdr: dns.RR_Header{
			Name:   "example.com.",
			Rrtype: dns.TypeSOA,
			Class:  dns.ClassINET,
			Ttl:    3600,
		},
		Ns:      "ns1.example.com.",
		Mbox:    "admin.example.com.",
		Serial:  2025010101,
		Refresh: 3600,
		Retry:   600,
		Expire:  86400,
		Minttl:  300,
	},

	// zone 内容
	&dns.A{
		Hdr: dns.RR_Header{
			Name:   "example.com.",
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    300,
		},
		A: net.ParseIP("1.2.3.4"),
	},

	&dns.A{
		Hdr: dns.RR_Header{
			Name:   "www.example.com.",
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    300,
		},
		A: net.ParseIP("5.6.7.8"),
	},
}

// SOA 结尾（AXFR 必须以同 SOA 记录结束）
var endSOA = &dns.SOA{
	Hdr: dns.RR_Header{
		Name:   "example.com.",
		Rrtype: dns.TypeSOA,
		Class:  dns.ClassINET,
		Ttl:    3600,
	},
	Ns:      "ns1.example.com.",
	Mbox:    "admin.example.com.",
	Serial:  2025010101,
	Refresh: 3600,
	Retry:   600,
	Expire:  86400,
	Minttl:  300,
}

type axfrHandler struct{}

func (h *axfrHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]

	// 仅允许 TCP 执行 AXFR
	if q.Qtype == dns.TypeAXFR && w.RemoteAddr().Network() == "tcp" {
		log.Println("收到 AXFR 请求")
		h.handleAXFR(w, r)
		return
	}

	// 其他请求（普通 A / AAAA 等）
	msg := &dns.Msg{}
	msg.SetReply(r)

	msg.Answer = append(msg.Answer, &dns.A{
		Hdr: dns.RR_Header{
			Name:   q.Name,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    60,
		},
		A: net.ParseIP("127.0.0.1"),
	})

	_ = w.WriteMsg(msg)
}

func (h *axfrHandler) handleAXFR(w dns.ResponseWriter, r *dns.Msg) {
	zoneName := r.Question[0].Name
	t := new(dns.Transfer)

	c := make(chan *dns.Envelope)
	err := t.Out(w, r, c)
	if err != nil {
		log.Println("AXFR Out error:", err)
		return
	}

	// 1. 发送起始 SOA
	c <- zoneRecords[0]
	if err := c.Send(zoneRecords[0]); err != nil {
		log.Println("AXFR Send error:", err)
		return
	}

	// 2. 发送所有记录
	for _, rr := range zoneRecords[1:] {
		if err := c.Send(rr); err != nil {
			log.Println("AXFR Send error:", err)
			return
		}
	}

	// 3. 发送结束 SOA
	if err := c.Send(endSOA); err != nil {
		log.Println("AXFR End SOA error:", err)
		return
	}

	log.Printf("AXFR 完成 zone = %s\n", zoneName)
}

func main() {
	handler := &axfrHandler{}

	tcpServer := &dns.Server{
		Addr:         ":53",
		Net:          "tcp",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	udpServer := &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: handler,
	}

	go func() {
		log.Println("启动 TCP DNS (AXFR 支持) :53")
		log.Fatal(tcpServer.ListenAndServe())
	}()

	log.Println("启动 UDP DNS :53")
	log.Fatal(udpServer.ListenAndServe())
}
