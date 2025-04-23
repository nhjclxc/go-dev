package main

import (
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
)

func main() {
	goRoute2 := rabbitmq.NewRabbitMQRouting("go", "go-Route2")
	goRoute2.RecieveRouting()
}
