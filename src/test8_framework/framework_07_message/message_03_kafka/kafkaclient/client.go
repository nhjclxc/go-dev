package kafkaclient

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func NewDialer() *kafka.Dialer {
	return &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		// SASL / TLS 以后统一在这里加
	}
}
