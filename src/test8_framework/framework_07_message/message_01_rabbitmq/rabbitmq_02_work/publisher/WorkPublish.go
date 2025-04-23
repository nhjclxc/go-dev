package main

import (
	"fmt"
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("" + "go-work")

	for i := 0; i <= 23; i++ {
		rabbitmq.PublishSimple("Hello go-work!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
