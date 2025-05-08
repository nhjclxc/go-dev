package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"order-service/internal/config"
	rpcUser "order-service/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	UserService rpcUser.UserService                                          // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserService: rpcUser.NewUserService(zrpc.MustNewClient(c.UserService)), // 手动代码
	}
}
