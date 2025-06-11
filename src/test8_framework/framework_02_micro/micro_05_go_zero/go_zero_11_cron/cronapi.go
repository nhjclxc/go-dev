package main

import (
	"flag"
	"fmt"
	"go_zero_11_cron/internal/logic"

	"go_zero_11_cron/internal/config"
	"go_zero_11_cron/internal/handler"
	"go_zero_11_cron/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/cronapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)


	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)


	// 注册定时任务
	ctx.Cron.AddFunc("@every 3s", func() {
		fmt.Println("每 30 秒执行一次任务")
		logic.CornDoSomething(ctx)
	})

	// 启动定时器
	ctx.Cron.Start()
	defer ctx.Cron.Stop()



	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
