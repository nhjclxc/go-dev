package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "protobuf_05_tls/proto"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 没有证书的连接方式
	//conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// could not greet: rpc error: code = Unavailable desc = connection error: desc = "error reading server preface: EOF"

	// 有证书的连接方式
	creds, _ := credentials.NewClientTLSFromFile("./server.crt", "")
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTlsServiceClient(conn)

	r, err := client.TlsInterface(ctx, &pb.TlsRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Response:", r.Message)

}
