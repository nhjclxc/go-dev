package kafkaclient

import (
	"context"
	"fmt"
	"message_03_kafka/kafkaclient"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	producer := kafkaclient.NewProducer(
		kafkaclient.Config{
			Brokers:      []string{"192.168.203.182:9092"},
			BatchSize:    10,
			BatchTimeout: time.Second,
			WriteTimeout: 5 * time.Second,
		},
		"nginx_error",
	)
	defer producer.Close()

	_ = producer.Send(context.Background(), nil, []byte("hello kafka"))

}
func TestConsumer(t *testing.T) {
	consumer := kafkaclient.NewConsumer(
		kafkaclient.Config{
			Brokers: []string{"192.168.203.182:9092"},
			//GroupID: "nginx-error-consumer",
		},
		"nginx_error",
	)
	defer consumer.Close()

	fmt.Println("开启消息读取...")
	for {
		msg, err := consumer.Read(context.Background())
		if err != nil {
			panic(err)
		}
		println(string(msg.Value))
	}

}
