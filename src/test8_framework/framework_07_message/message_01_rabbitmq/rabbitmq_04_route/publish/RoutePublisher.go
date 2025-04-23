package main

import (
	"fmt"
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
	"strconv"
	"time"
)

// 发布者
func main() {

	route1 := rabbitmq.NewRabbitMQRouting("go", "go-Route1")
	route2 := rabbitmq.NewRabbitMQRouting("go", "go-Route2")
	for i := 0; i <= 13; i++ {
		route1.PublishRouting("Hello go-Route1 one!" + strconv.Itoa(i))
		route2.PublishRouting("Hello go-Route2 Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
