package kafkaclient

//
//import (
//	"testing"
//	"time"
//)
//
//func Test111(t *testing.T) {
//	producer := kafkaclient.NewProducer(
//		kafkaclient.Config{
//			Brokers:      []string{"192.168.203.182:9092"},
//			BatchSize:    10,
//			BatchTimeout: time.Second,
//			WriteTimeout: 5 * time.Second,
//		},
//		"nginx_error",
//	)
//
//	defer producer.Close()
//
//	_ = producer.Send(context.Background(), nil, []byte("hello kafka"))
//
//}
