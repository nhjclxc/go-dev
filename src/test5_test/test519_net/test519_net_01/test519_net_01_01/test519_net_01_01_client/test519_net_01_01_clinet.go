package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	//
	/*
	   type Conn interface {
	       // Read从连接中读取数据
	       // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	       Read(b []byte) (n int, err error)

	       // Write从连接中写入数据
	       // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	       Write(b []byte) (n int, err error)

	       // Close方法关闭该连接
	       // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
	       Close() error

	       // 返回本地网络地址
	       LocalAddr() Addr

	       // 返回远端网络地址
	       RemoteAddr() Addr

	       // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
	       // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
	       // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
	       // 参数t为零值表示不设置期限
	       SetDeadline(t time.Time) error

	       // 设定该连接的读操作deadline，参数t为零值表示不设置期限
	       SetReadDeadline(t time.Time) error

	       // 设定该连接的写操作deadline，参数t为零值表示不设置期限
	       // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
	       SetWriteDeadline(t time.Time) error
	   }
	*/

	// 客户端通过 net.Dial() 创建了一个和服务器之间的连接。
	//当然，服务器必须先启动好，如果服务器并未开始监听，客户端是无法成功连接的。
	// 如果在服务器没有开始监听的情况下运行客户端程序，客户端会停止并打印出以下错误信息：
	// dial err dial tcp [::1]:50000: connectex: No connection could be made because the target machine actively refused it.

	// 在网络编程中 net.Dial() 函数是非常重要的，一旦你连接到远程系统，函数就会返回一个 Conn 类型的接口，我们可以用它发送和接收数据。
	// Dial() 函数简洁地抽象了网络层和传输层。所以不管是 IPv4 还是 IPv6，TCP 或者 UDP 都可以使用这个公用接口。

	connTimeout := time.Now().Add(3 * time.Second)
	fmt.Println("connTimeout:", connTimeout)

	// 连接tcp服务器
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("dial err", err)
		return
	}

	fmt.Println(conn)
	fmt.Println("LocalAddr: ", conn.LocalAddr())
	fmt.Println("RemoteAddr: ", conn.RemoteAddr())
	conn.SetDeadline(connTimeout) // 连接到什么时候超时，并不是多久以后超时

	// 发送一个数据试试
	write, errWrite := conn.Write([]byte("Ping"))
	if errWrite != nil {
		fmt.Println("errWrite: ", errWrite)
		return
	}
	fmt.Println("write: ", write)

	// 读取一个数据看看
	var buf []byte = make([]byte, 512)
	read, err := conn.Read(buf)
	if err != nil {
		return
	}
	fmt.Println("buf: ", string(buf))

	log.Println("read = ", read, "buf = ", string(buf))

	for true {
		// 不断的进行消息的收发操作
		// ......
		break
	}

	conn.Close()

}
