package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_order/internal/server"

	"grpc_order/grpc/order"
	"grpc_order/internal/config"
	"grpc_order/internal/handler"
	"grpc_order/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {

	// 加载配置
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)


	// 启动grpc服务
	go func() {
		grpcServer := zrpc.MustNewServer(c.RpcConf, func(grpcServer *grpc.Server) {
			order.RegisterGrpcOrderServiceServer(grpcServer, server.NewGrpcOrderServiceServer(ctx))

			if c.RpcConf.Mode == service.DevMode || c.RpcConf.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		defer grpcServer.Stop()

		fmt.Printf("Starting rpc server at %s...\n", c.RpcConf.ListenOn)
		grpcServer.Start()

	}()


	// 启动http服务
	httpServer := rest.MustNewServer(c.RestConf)
	defer httpServer.Stop()

	handler.RegisterHandlers(httpServer, ctx)

	fmt.Printf("Starting http server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	httpServer.Start()

}
