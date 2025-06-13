package main

import (
	"flag"
	"fmt"
	rabbitMQMessageHandler "go_zero_15_rabbitmq/internal/rabbitmq"

	"go_zero_15_rabbitmq/internal/config"
	"go_zero_15_rabbitmq/internal/handler"
	"go_zero_15_rabbitmq/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/rabbitmqApi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)


	// 注册rabbitmq的消费者
	go rabbitMQMessageHandler.NewRabbitMqMessageHandler(ctx.RabbitMQ).MessageHandler()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
