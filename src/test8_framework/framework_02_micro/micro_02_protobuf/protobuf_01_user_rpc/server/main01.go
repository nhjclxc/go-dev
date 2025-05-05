package main
//
//import (
//	"fmt"
//	"net"
//	"net/rpc"
//	dto "server/dto"
//)
//
//// 定义接口结构体
//type UserService struct {}
//
//func (this *UserService) GetUserByName(request *dto.UserRequest, response *dto.UserResponse) error {
//
//	fmt.Printf("服务端接收到的数据 %#v \n", request)
//
//
//	*response = dto.UserResponse{
//		Id: 666,
//		Name: "我是" + request.GetName(),
//		Age: 18,
//		Hobby: []string{"打飞机", "开坦克", "调情"},
//	}
//
//	return nil
//}
//
//
//func main01() {
//
//
//	// 注册服务
//	err := rpc.RegisterName("UserService", new(UserService))
//	if err != nil {
//		return
//	}
//
//	// 监听端口
//	listener, err := net.Listen("tcp", "127.0.0.1:8090")
//	if err != nil {
//		return
//	}
//	defer listener.Close()
//
//	// 等待客户端连接
//	for true {
//		conn, err := listener.Accept()
//		if err != nil {
//			return
//		}
//
//		// 开一个协程去处理
//		go rpc.ServeConn(conn)
//	}
//
//}
