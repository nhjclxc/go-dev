package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_user/grpc/user"
	"grpc_user/internal/handler"
	"grpc_user/internal/server"

	"grpc_user/internal/config"
	"grpc_user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/anonymous_user.yaml", "the config file")

func main() {

	// 加载配置
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 启动grpc服务
	go func() {
		grpcServer := zrpc.MustNewServer(c.RpcConf, func(grpcServer *grpc.Server) {
			user.RegisterGrpcUserServiceServer(grpcServer, server.NewGrpcUserServiceServer(ctx))

			if c.RpcConf.Mode == service.DevMode || c.RpcConf.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		defer grpcServer.Stop()

		fmt.Printf("Starting rpc server at %s...\n", c.RpcConf.ListenOn)
		grpcServer.Start()

		// 查看当前 anonymous_user.rpc 是否注册到了etcd里面
		// [root@iZ2zehmay1c73eaheydrpyZ /]# docker exec etcd etcdctl --endpoints=39.106.59.225:2379 get anonymous_user.rpc --prefix
		//anonymous_user.rpc/3368577425648289896
		//192.168.8.131:9190

	}()

	// 启动http服务
	httpServer := rest.MustNewServer(c.RestConf)
	defer httpServer.Stop()

	handler.RegisterHandlers(httpServer, ctx)

	fmt.Printf("Starting http server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	httpServer.Start()
}
