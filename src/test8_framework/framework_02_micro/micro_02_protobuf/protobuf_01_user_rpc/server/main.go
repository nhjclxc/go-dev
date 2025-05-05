package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	grpc_userinfo "server/grpc/userinfo"
)

// 定义接口结构体
type UserInfoService struct {
	// 继承 proto 生成的一个接口
	grpc_userinfo.UserInfoServiceServer
}

var userInfoService = UserInfoService{}

// GetUserInfoByName 更具名族查询用户
// 注意方法的参数和返回值，和没有使用protobuf的时候略有不同
func (this *UserInfoService) GetUserInfoByName(context context.Context, request *grpc_userinfo.UserRequest) (*grpc_userinfo.UserResponse, error) {

	fmt.Printf("服务端接收到的数据 %#v \n", request)

	response := grpc_userinfo.UserResponse{
		Id: 666,
		Name: "我是" + request.GetName(),
		Age: 18,
		Hobby: []string{"打飞机", "开坦克", "调情"},
	}

	return &response, nil
}



// go get google.golang.org/grpc
// go get github.com/golang/protobuf/proto
func main() {
	// 1.监听端口
	// 2.需要实例化gRPC服务端
	// 3.在gRPC商注册微服务
	// 4.启动服务端



	// 1.监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
		return
	}

	// 2.实例化gRPC
	server := grpc.NewServer()

	// 3.在gRPC上注册微服务
	grpc_userinfo.RegisterUserInfoServiceServer(server, &userInfoService)

	// 启用反射，方便调试
	reflection.Register(server)

	// 4.启动服务器
	fmt.Println("Server is running on port 8090...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}
