package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
	pb "protobuf_03_test1/proto"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUser 实现grpc接口
func (s server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.CommonResponse, error) {

	fmt.Printf("GetUser.GetUserRequest = %#v \n", req)

	data, err := anypb.New(&pb.GetUserResponse{
		UserId:   666,
		Age:      18,
		UserName: "我是一个用户啊",
	})
	if err != nil {
		fmt.Printf("GetUser.GetUserRequest.anypb.New = %#v \n", err)
		return nil, err
	}
	res := &pb.CommonResponse{
		Code: 200,
		Msg:  "请求成功！",
		Data: data,
	}

	return res, nil
}

func main() {
	// 启动端口监听
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("grpc服务端启动失败，50051端口监听失败！ err = %v \n", err.Error())
		return
	}

	// 创建rpc服务
	rpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(rpcServer, &server{})
	fmt.Printf("grpc服务启动成功，监听50051端口")
	if err := rpcServer.Serve(listen); err != nil {
		fmt.Println("grpc服务启动失败，", err)
		return
	}

}
