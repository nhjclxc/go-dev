package pkg01_dsn

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"testing"
)

// go get github.com/miekg/dns
// github.com/miekg/dns（简称 miekg/dns）是一个非常强大的 Go 语言 DNS 库。下面我详细讲讲这个包（library）能做什么、它的核心功能、设计哲学，以及典型使用场景。

func Test11(t *testing.T) {

	// 创建一个 DNS 消息
	m := new(dns.Msg)
	// 设置要查询的问题 (这里查 A 记录)
	m.SetQuestion(dns.Fqdn("google.com"), dns.TypeA)
	// 递归查询
	m.RecursionDesired = true

	// 创建一个客户端
	c := new(dns.Client)
	// 向 DNS 服务器发送请求 (假设使用 Google DNS)
	in, rtt, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}

	fmt.Printf("RTT: %v\n", rtt)
	// 遍历返回的 Answer 部分
	for _, ans := range in.Answer {
		// 类型断言为 A 记录
		if a, ok := ans.(*dns.A); ok {
			fmt.Printf("A: %s -> %s\n", a.Hdr.Name, a.A.String())
		} else {
			fmt.Printf("其他记录: %v\n", ans)
		}
	}

}

func Test22(t *testing.T) {
	// 简单的dns查询使用；net.LookupIP
	domain := "google.com"

	ip, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("lookup ip err: %s \n", err)
		return
	}

	fmt.Printf("%s ---ip---> %s \n", domain, ip)

}
