package main

import (
	"fmt"
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
	"strconv"
	"time"
)

// 发布者
func main() {
	rabbitmq := rabbitmq.NewRabbitMQPubSub("" + "go-PubSub")
	for i := 0; i <= 13; i++ {
		rabbitmq.PublishPub("订阅模式生产第" + strconv.Itoa(i) + "条" + "数据")
		fmt.Println("订阅模式生产第" + strconv.Itoa(i) + "条" + "数据")
		time.Sleep(1 * time.Second)
	}
}
