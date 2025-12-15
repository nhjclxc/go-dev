package kafkaclient

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(cfg Config, topic string) *Consumer {
	readerCfg := kafka.ReaderConfig{
		Brokers:        cfg.Brokers,
		Topic:          topic,
		Dialer:         NewDialer(),
		SessionTimeout: cfg.ReadTimeout,
		//ClientID:    cfg.ClientID,
	}

	if cfg.GroupID != "" {
		readerCfg.GroupID = cfg.GroupID
	} else {
		readerCfg.StartOffset = cfg.StartOffset
	}

	reader := kafka.NewReader(readerCfg)

	return &Consumer{reader: reader}
}

func (c *Consumer) Read(ctx context.Context) (kafka.Message, error) {
	return c.reader.ReadMessage(ctx)
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
