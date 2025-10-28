package main

import (
	"fmt"
	"net"
	"testing"
)

// 解析域名的CNAME
func TestName(t *testing.T) {

	name := "www.baidu.com"
	name = ""

	// 找到它的 CNAME（如果存在，返回最终的别名）
	cname, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("LookupCNAME error:", err)
		return
	}
	fmt.Println("CNAME:", cname)

	// 再查最终的 A/AAAA
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("LookupHost error:", err)
		return
	}
	fmt.Println("A/AAAA:", addrs)

}
