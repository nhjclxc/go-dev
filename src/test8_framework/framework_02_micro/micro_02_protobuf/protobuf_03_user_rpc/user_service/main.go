package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	grpc_user "user_service/grpc/user"
)

// 定义 User 微服务结构体
type UserService struct {
	grpc_user.UserServiceServer
}

var userService UserService

func init() {
	userService = UserService{}
}

// GetUser 根据指定条件获取数据
func (this *UserService) GetUser(context context.Context, request *grpc_user.UserRequest) (*grpc_user.UserResponse, error) {

	fmt.Printf("UserService.GetUser：接收到调用者的请求数据：%#v 。\n", request)

	id := request.GetId()
	if id == 666 {
		// 表示找到数据了，返回指定数据，否则不返回数据
		response := grpc_user.UserResponse{
			Id:      666,
			Name:    "无名者",
			IdCard:  "123456789",
			Age:     18,
			Address: "不知道住在哪里",
			Hobby:   []string{"看片", "撸管", "打飞机"},
		}

		return &response, nil
	}

	return nil, errors.New("找不到指定数据！！！")

}

// go get google.golang.org/grpc
// go get github.com/golang/protobuf/proto
func main() {
	// 1.监听端口
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		return
	}
	defer listener.Close()


	// 2.实例化gRPC服务端
	server := grpc.NewServer()

	// 3.在gRPC里注册微服务
	grpc_user.RegisterUserServiceServer(server, &userService)


	// 4.启动服务端
	err2 := server.Serve(listener)
	if err2 != nil {
		fmt.Printf("UserService服务注册失败！！！\n", err2)
		return
	}
	fmt.Printf("UserService is running on port 8090...\n")

}
