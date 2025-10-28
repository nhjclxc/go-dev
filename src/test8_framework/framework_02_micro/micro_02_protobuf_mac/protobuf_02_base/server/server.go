package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "protobuf_02_base/proto"
)

type server struct {
	pb.UnimplementedBaseServiceServer
}

func (s *server) GetBase(ctx context.Context, req *pb.GetBaseRequest) (*pb.GetBaseResponse, error) {
	fmt.Printf("接收[%s]的到请求 \n", req.Name)
	return &pb.GetBaseResponse{Res: "Res " + req.Name, Data: "Data " + req.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051") // gRPC 服务端口
	if err != nil {
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBaseServiceServer(grpcServer, &server{})
	fmt.Println("🚀 gRPC server listening on :50051")
	grpcServer.Serve(lis)
}
