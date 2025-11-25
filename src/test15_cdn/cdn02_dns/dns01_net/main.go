package main

import (
	"fmt"
	"log"
	"net"
)

// 使用go语言的net原生包实现dns服务器解析，将local.com解析为127.0.0.1
func main() {

	// 1、创建dns服务器，dns常用的端口是：53，常用的协议是：udp
	// dns连接信息
	addr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"), // 表示对所有ip生效
		Port: 53,
	}
	// 创建dns监听
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("dns连接53端口监听失败！！！")
		return
	}
	defer conn.Close()
	log.Println("DNS server running on UDP :53 ...")

}
