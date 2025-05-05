package main
//
//import (
//	dto "client/dto"
//	"fmt"
//	"net"
//	"net/rpc"
//)
//
//func main() {
//
//
//	// 连接服务器
//	conn, err := net.Dial("tcp", "127.0.0.1:8090")
//	if err != nil {
//		return
//	}
//	defer conn.Close()
//
//	// 注册客户端
//	cleint := rpc.NewClient(conn)
//
//	request := dto.UserRequest{Name: "张三"}
//	response := dto.UserResponse{}
//
//
//	// 调用远程接口
//	err2 := cleint.Call("UserInfoService.GetUserByName", &request, &response)
//	if err2 != nil {
//		return
//	}
//
//	fmt.Printf("接收到服务器的数据：%#v \n", response)
//
//
//}
