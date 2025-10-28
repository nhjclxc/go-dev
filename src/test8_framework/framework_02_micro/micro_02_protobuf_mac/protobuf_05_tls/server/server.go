package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	pb "protobuf_05_tls/proto"
)

type server struct {
	pb.UnimplementedTlsServiceServer
}

func (s *server) TlsInterface(ctx context.Context, req *pb.TlsRequest) (*pb.TlsResponse, error) {
	fmt.Printf("接收[%s]的到请求 \n", req.Name)
	return &pb.TlsResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // gRPC 服务端口
	if err != nil {
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}
	// 这个是没有证书的服务端
	//grpcServer := grpc.NewServer()

	// 读取证书的服务端
	creds, _ := credentials.NewServerTLSFromFile("./server.crt", "./server.key")
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterTlsServiceServer(grpcServer, &server{})
	fmt.Println("🚀 gRPC server listening on :50051")
	grpcServer.Serve(lis)
}
