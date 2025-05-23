package main

import (
	"flag"
	"fmt"

	"zero-order/internal/config"
	"zero-order/internal/handler"
	"zero-order/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order-api.yaml", "the config file")

// cmd启动命令：go run order.go -f etc/order-api.yaml
// 访问接口：curl -X POST http://127.0.0.1:8090/order/info -H "Content-Type: application/json" -d "{\"orderId\":666}"
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
}
