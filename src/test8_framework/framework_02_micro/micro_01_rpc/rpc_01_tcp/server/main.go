package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义微服务类对象
type MicroRpcTcp struct {
	Name string
}

// 绑定微服务方法
func (this *MicroRpcTcp) HelloWorld(requst string, response *string) error {
	fmt.Println("接收到了请求：", requst)
	*response  = requst + "你好世界，HelloWorld"
	fmt.Println("处理了请求：", *response)
	return nil
}


func main() {

	// 1、注册 RPC 服务，MicroRpcTcpServer 是服务名，在客户端调用的时候使用的就是这个
	err := rpc.RegisterName("MicroRpcTcpServer", new(MicroRpcTcp))
	if err != nil {
		fmt.Println("微服务注册失败！！！", err)
		return
	}

	// 2、设置监听
	listener, err2 := net.Listen("tcp", "127.0.0.1:8090")
	if err2 != nil {
		fmt.Println("微服务监听失败！！！", err2)
		return
	}
	// 服务关闭之后，释放端口
	defer listener.Close()


	// 3、接受客户端连接
	for true {
		// 等待连接
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println("微服务等待客户端连接失败！！！", err3)
			return
		}

		// 开一个协程处理具体的业务
		go rpc.ServeConn(conn)
	}

}
