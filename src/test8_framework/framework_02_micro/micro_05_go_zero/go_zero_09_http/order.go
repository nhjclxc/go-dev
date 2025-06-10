package main

import (
	"flag"
	"fmt"


	"go_zero_09_http/internal/config"
	"go_zero_09_http/internal/handler"
	"go_zero_09_http/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

)

var configFile = flag.String("f", "etc/order-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()


	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)


	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()



	// 统一响应结构体包装
	// https://go-zero.dev/docs/tutorials/http/server/response/ext
	//errors.CodeMsg{}
	//"github.com/zeromicro/x/errors"


}

/*
在 go-zero 中，支持通过 http.Request 的 Body 字段获取请求参数，同时支持通过 http.Request 的 Form 字段获取请求参数，
其中 Body 只接受 application/json 格式的请求参数，Form 只接受 application/x-www-form-urlencoded 格式的请求参数，
除此之外，go-zero 还支持 path 参数和请求头参数的获取，他们都是通过 httpx.Parse 方法将数据解析到结构体中。

MiddlewaresConf



*/