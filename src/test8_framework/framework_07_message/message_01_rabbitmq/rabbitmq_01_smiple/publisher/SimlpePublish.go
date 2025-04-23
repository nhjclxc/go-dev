package main

import (
	"fmt"
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("" + "go-simple")
	rabbitmq.PublishSimple("Hello go-simple!")
	fmt.Println("发送成功！")
}
