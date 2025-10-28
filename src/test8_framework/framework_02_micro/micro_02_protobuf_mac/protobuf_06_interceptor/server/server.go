package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "protobuf_06_interceptor/proto"
	"time"
)

type server struct {
	pb.UnimplementedInterceptorServiceServer
}

func (s *server) InterceptorInterface(ctx context.Context, req *pb.InterceptorRequest) (*pb.InterceptorResponse, error) {
	fmt.Printf("接收[%s]的到请求 \n", req.Name)
	return &pb.InterceptorResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // gRPC 服务端口
	if err != nil {
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(serverInterceptor), // 定义服务端拦截器
	)
	pb.RegisterInterceptorServiceServer(grpcServer, &server{})
	fmt.Println("🚀 gRPC server listening on :50051")
	grpcServer.Serve(lis)
}

func serverInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {

	start := time.Now()
	log.Printf("[Server Interceptor] %s called", info.FullMethod)

	// 执行业务逻辑
	resp, err = handler(ctx, req)

	log.Printf("[Server Interceptor] Completed in %v, err=%v", time.Since(start), err)
	return resp, err
}
