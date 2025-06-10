package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc_coupon/internal/config"

	"grpc_coupon/grpc/grpcuserservice"
)

type ServiceContext struct {
	Config config.Config
	GrpcUserService grpcuserservice.GrpcUserService  // 手动代码, 建立rpc连接
}

func NewServiceContext(c config.Config) *ServiceContext {

	grpcUserClient, err := zrpc.NewClient(c.RpcUserConf)
	if err != nil {
		//return nil
	}

	return &ServiceContext{
		Config: c,
		// zrpc.MustNewClient是直连模式，是一个阻塞的的gpc连接方式，回影响本服务的启动
		// zrpc.NewClient是懒加载连接模式，是非阻塞的，不会影响本服务的启动

		GrpcUserService: grpcuserservice.NewGrpcUserService(grpcUserClient), // 手动代码, 建立rpc连接
	}
}