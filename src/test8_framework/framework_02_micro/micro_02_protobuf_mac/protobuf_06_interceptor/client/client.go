package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "protobuf_06_interceptor/proto"
)

func clientInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	log.Printf("[Client Interceptor] Calling method: %s", method)

	// 调用实际RPC
	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Printf("[Client Interceptor] Done: %v, Duration: %s", err, time.Since(start))
	return err
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(clientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewInterceptorServiceClient(conn)

	r, err := client.InterceptorInterface(ctx, &pb.InterceptorRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Response:", r.Message)

}
