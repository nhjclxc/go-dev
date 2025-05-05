package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Http01 struct {
}

func (this *Http01) GetById(request string, response *string) error {

	fmt.Println("Http01.GetById = ", request)

	*response = "http服务端：" + request;

	fmt.Println("Http01.GetById.*response = ", *response)

	return nil
}


func main() {


	// 注册服务
	err := rpc.RegisterName("Http01", new(Http01))
	if err != nil {
		return
	}

	listener, err := net.Listen("http", "127.0.0.1:8090")
	if err != nil {
		return
	}
	defer listener.Close()

	for true {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		// 开一个协程去处理这个请求
		go rpc.ServeConn(conn)
	}


}
