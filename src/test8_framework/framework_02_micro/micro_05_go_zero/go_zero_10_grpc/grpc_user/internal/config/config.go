package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RestConf     rest.RestConf      `json:"restConf"`     // HTTP 配置
	RpcConf      zrpc.RpcServerConf `json:"RpcConf"`      // GRPC 配置
}
