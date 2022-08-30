package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	w *kafka.Writer
}

func New(c *Config) *Kafka {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(c.Brokers...),
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		MaxAttempts:            c.MaxAttempts,
		Async:                  c.Async,
	}
	return &Kafka{
		w: w,
	}
}

func (kf *Kafka) Write(topic string, k []byte, v []byte) error {
	err := kf.w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   k,
			Value: v,
			Topic: topic,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (kf *Kafka) Close() error {
	return kf.w.Close()
}
