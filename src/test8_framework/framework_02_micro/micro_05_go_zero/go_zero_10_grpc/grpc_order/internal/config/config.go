package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

// 官方配置示例：https://go-zero.dev/docs/tutorials/grpc/client/configuration
type Config struct {
	RestConf     rest.RestConf      `json:"restConf"`     // HTTP 配置
	RpcConf      zrpc.RpcServerConf `json:"RpcConf"`      // GRPC 配置
	RpcCouponConf zrpc.RpcClientConf `json:"RpcCouponConf"` // GRPC Order 的配置
	RpcUserConf  zrpc.RpcClientConf `json:"RpcUserConf"`  // GRPC User 的配置
}

// RpcClientConf
//名称	类型	含义	默认值	是否必选
//Etcd	EtcdConf	服务发现配置，当需要使用 etcd 做服务发现时配置	无	否
//Endpoints	string 类型数组	RPC Server 地址列表，用于直连，当需要直连 rpc server 集群时配置	无	否
//Target	string	域名解析地址，名称规则请参考 https://github.com/grpc/grpc/blob/master/doc/naming.md	无	否
//App	string	rpc 认证的 app 名称，仅当 rpc server 开启认证时配置	无	否
//Token	string	rpc 认证的 token，仅当 rpc server 开启认证时配置	无	否
//NonBlock	bool	是否阻塞模式,当值为 true 时，不会阻塞 rpc 链接	false	否
//Timeout	int64	超时时间	2000ms	否
//KeepaliveTime	Time.Duration	保活时间	20s	否
//Middlewares	ClientMiddlewaresConf	是否启用中间件

type RpcClientConf struct {
	Etcd          discov.EtcdConf `json:",optional,inherit"`
	Endpoints     []string        `json:",optional"`
	Target        string          `json:",optional"`
	App           string          `json:",optional"`
	Token         string          `json:",optional"`
	NonBlock      bool            `json:",optional"`
	Timeout       int64           `json:",default=2000"`
	KeepaliveTime time.Duration   `json:",default=20s"`
	//Middlewares   ClientMiddlewaresConf
}

type EtcdConf struct {
	Hosts              []string
	Key                string
	ID                 int64  `json:",optional"`
	User               string `json:",optional"`
	Pass               string `json:",optional"`
	CertFile           string `json:",optional"`
	CertKeyFile        string `json:",optional=CertFile"`
	CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool   `json:",optional"`
}

type ServerMiddlewaresConf struct {
	Trace      bool `json:",default=true"`
	Recover    bool `json:",default=true"`
	Stat       bool `json:",default=true"`
	Prometheus bool `json:",default=true"`
	Breaker    bool `json:",default=true"`
}
