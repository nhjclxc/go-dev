package main

import (
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
)

func main() {
	goRoute1 := rabbitmq.NewRabbitMQRouting("go", "go-Route1")
	goRoute1.RecieveRouting()
}
