package kafkaclient

import "time"

type Config struct {
	Brokers []string

	// Producer
	WriteTimeout time.Duration
	BatchSize    int
	BatchTimeout time.Duration

	// Consumer
	GroupID     string
	StartOffset int64
	ReadTimeout time.Duration

	// 通用
	ClientID string
}
