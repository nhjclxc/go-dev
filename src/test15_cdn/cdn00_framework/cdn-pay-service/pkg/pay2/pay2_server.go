package main

import (
	"context"
	"fmt"
	pb "github.com/yourorg/cdn-common/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedPay2ServiceServer
}

func (s *server) Charge(ctx context.Context, req *pb.PayChargeRequest) (*pb.PayChargeResponse, error) {
	fmt.Printf("Charge 接收[%s]的到请求 \n", req.UserID)
	return &pb.PayChargeResponse{Message: "Charge " + req.UserID}, nil
}
func (s *server) QueryStatus(ctx context.Context, req *pb.PayQueryStatusRequest) (*pb.PayQueryStatusResponse, error) {
	fmt.Printf("QueryStatus 接收[%s]的到请求 \n", req.OrderID)
	return &pb.PayQueryStatusResponse{Message: "QueryStatus " + req.OrderID}, nil
}

// 即提供了grpc服务，又提供了http服务
func main() {

	lis, err := net.Listen("tcp", ":50051") // gRPC 服务端口
	if err != nil {
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPay2ServiceServer(grpcServer, &server{})
	fmt.Println("🚀 gRPC server listening on :50051")
	grpcServer.Serve(lis)

}
