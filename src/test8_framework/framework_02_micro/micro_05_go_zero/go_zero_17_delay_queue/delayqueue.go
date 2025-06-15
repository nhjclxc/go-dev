package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_17_delay_queue/internal/config"
	"go_zero_17_delay_queue/internal/handler"
	"go_zero_17_delay_queue/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/delayQueue.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)


	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)





	ctx.DqConsumer.Consume(func(body []byte) {
		logx.Infof("consumer job  %s \n", string(body))
	})


	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
