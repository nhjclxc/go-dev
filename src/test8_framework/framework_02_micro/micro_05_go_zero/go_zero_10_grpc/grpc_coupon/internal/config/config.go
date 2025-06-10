package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RestConf     rest.RestConf      `json:"restConf"`     // HTTP 配置
	RpcConf      zrpc.RpcServerConf `json:"RpcConf"`      // GRPC 配置
	RpcOrderConf zrpc.RpcClientConf `json:"RpcOrderConf"` // GRPC Order 的配置
	RpcUserConf  zrpc.RpcClientConf `json:"RpcUserConf"`  // GRPC User 的配置
}
