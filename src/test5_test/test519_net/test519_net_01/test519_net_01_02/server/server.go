package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// 1、开启一个服务器
	listen, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatalln("server.listen.err: ", err.Error())
		return
	}

	// 2、监听接受客户端的连接
	for true {
		accept, errAccept := listen.Accept()
		if errAccept != nil {
			log.Fatalln("server.accept.err: ", errAccept.Error())
			return
		}

		// 3、开启一个协程去处理这个客户端的请求
		go dealClientRequest(accept)
	}

}

// 处理来自客户端的请求
func dealClientRequest(accept net.Conn) {
	fmt.Println("LocalAddr: ", accept.LocalAddr())
	fmt.Println("RemoteAddr: ", accept.RemoteAddr())

	for true {
		// 4、等待客户端发送来消息
		buf := make([]byte, 512)
		// accept.Read 会阻塞等待客户端的数据
		read, errAccept := accept.Read(buf)
		if errAccept != nil {
			log.Fatalln("server.dealClientRequest.err: ", errAccept.Error())
			return
		}
		log.Println("server.dealClientRequest.read: size = ", read, ", read = ", string(buf))

		if "CLOSE_SERVER" == string(buf) {
			accept.Close()
			log.Println("server.CLOSE_SERVER")
			return
		}

		// 响应数据
		_, errWrite := accept.Write([]byte("Pong-" + string(buf)))
		if errWrite != nil {
			log.Fatalln("server.errWrite.err: ", errWrite.Error())
			return
		}
	}

}
