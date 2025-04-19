package main

import (
	"fmt"
	"net"
)

func main() {

	/*
		这部分我们将使用 TCP 协议和在 14 章讲到的协程范式编写一个简单的客户端-服务器应用，一个 (web) 服务器应用需要响应众多客户端的并发请求：
		Go 会为每一个客户端产生一个协程用来处理请求。我们需要使用 net 包中网络通信的功能。它包含了处理 TCP/IP 以及 UDP 协议、域名解析等方法。
	*/

	/*
		net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。

		虽然本包提供了对网络原语的访问，大部分使用者只需要Dial、Listen和Accept函数提供的基本接口；以及相关的Conn和Listener接口。crypto/tls包提供了相同的接口和类似的Dial和Listen函数。
	*/

	/**
	在 main() 中创建了一个 net.Listener 类型的变量 listener，他实现了服务器的基本功能：
	用来监听和接收来自客户端的请求（基于 TCP 协议下，位于 IP 地址为 127.0.0.1、端口为 50000 的 localhost）。
	Listen() 函数可以返回一个 error 类型的错误变量。用一个无限 for 循环的 listener.Accept() 来等待客户端的请求。
	客户端的请求将产生一个 net.Conn 类型的连接变量。然后一个独立的协程使用这个连接执行 doServerStuff()，
	开始使用一个 512 字节的缓冲 data 来读取客户端发送来的数据，并且把它们打印到服务器的终端，
	len() 获取客户端发送的数据字节数；当客户端发送的所有数据都被读取完成时，协程就结束了。
	这段程序会为每一个客户端连接创建一个独立的协程。必须先运行服务器代码，再运行客户端代码。
	*/
	// Listen函数创建的服务端：

	// https://studygolang.com/static/pkgdoc/pkg/net.htm#Listener

	fmt.Println("Starting te server ...50000")
	// 创建一个 Listener 来监听 50000 端口
	// 网络类型参数net必须是面向流的网络："tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
	// 对TCP和UDP网络，地址格式是host:port或[host]:port
	// func Listen(net, laddr string) (Listener, error)
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error Listening ...", err)
		return // 程序终止
	}

	fmt.Println(listener)

	// 监听并接受来自客户端的链接
	for {
		// func (l *TCPListener) Accept() (Conn, error)
		// Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。
		conn, connErr := listener.Accept()
		if connErr != nil {
			fmt.Println("Error Accepting ...", connErr)
			//return // 程序终止
			continue
		}
		fmt.Println(conn)

		// 每监听到一个客户端就开启一个协程来处理这个客户端
		go doServerStuff(conn)

	}

}

// 服务器处理客户端的链接
// conn net.Conn: 客户端连接对象
func doServerStuff(conn net.Conn) {
	fmt.Println("conn.LocalAddr()", conn.LocalAddr())
	fmt.Println("conn.RemoteAddr()", conn.RemoteAddr())

	// 不停的去读取客户端的数据
	for {
		var buf []byte = make([]byte, 512)
		// func (c *IPConn) Read(b []byte) (int, error)
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading ...", err)
			if "EOF" == err.Error() {
				fmt.Println("客户端关闭连接！！！")
			}
			return // 程序终止
		}
		fmt.Printf("Received: len = %d, data = %s", size, string(buf[:size]))

		if "Ping" == string(buf[:size]) {
			write, err := conn.Write([]byte("Pong"))
			if err != nil {
				return
			}
			fmt.Println(write)
		}

	}

}
