package main

import (
	"fmt"
	"go-dev/src/test8_framework/framework_07_message/message_01_rabbitmq/rabbitmq"
	"strconv"
	"time"
)

// 发布者
func main() {

	goTopic1 := rabbitmq.NewRabbitMQTopic("exchangeTopic", "go.topic.one")
	goTopic2 := rabbitmq.NewRabbitMQTopic("exchangeTopic", "go.topic.two")
	for i := 0; i <= 100; i++ {
		goTopic1.PublishTopic("Hello exchangeTopic topic one!" + strconv.Itoa(i))
		goTopic2.PublishTopic("Hello exchangeTopic topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
