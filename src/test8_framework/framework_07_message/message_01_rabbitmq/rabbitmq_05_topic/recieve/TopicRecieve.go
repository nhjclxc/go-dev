package main

import (
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
)

func main() {
	goTopic := rabbitmq.NewRabbitMQTopic("exchangeTopic", "*")
	goTopic.RecieveTopic()
}
