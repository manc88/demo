package messagebroker

import (
	userservice "github.com/manc88/demo/internal/user_service"
	"github.com/manc88/demo/pkg/kafka"
)

var _ userservice.IMessageBroker = (*MessageBroker)(nil)

type MessageBroker struct {
	k *kafka.Kafka
}

func New(k *kafka.Kafka) *MessageBroker {
	return &MessageBroker{
		k: k,
	}
}

func (m *MessageBroker) Write(dest string, data []byte) error {
	err := m.k.Write(dest, nil, data)
	return err
}

func (m *MessageBroker) Close() error {
	return m.k.Close()
}
