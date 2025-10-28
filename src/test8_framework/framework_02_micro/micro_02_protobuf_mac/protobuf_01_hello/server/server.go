package main

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"net"
//
//	"google.golang.org/grpc"
//	pb "protobuf_01_hello/proto"
//)
//
//type server struct {
//	pb.UnimplementedHelloServiceServer
//}
//
//func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
//	fmt.Printf("接收[%s]的到请求 \n", req.Name)
//	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
//}
//
//func (s *server) SayMorning(ctx context.Context, req *pb.MorningRequest) (*pb.MorningResponse, error) {
//	log.Printf("Morning call: %v", req.GetName())
//	return &pb.MorningResponse{Greeting: "Good morning, " + req.GetName() + "!"}, nil
//}
//
// // 只提供了grpc服务
//func main() {
//	lis, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterHelloServiceServer(s, &server{})
//	fmt.Println("🚀 gRPC server listening on :50051")
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
