// main_test.go
package main

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
)

// MockProducer is a mock implementation of the Producer interface
type MockProducer struct {
	messages []*sarama.ProducerMessage
}

// SendMessage mocks sending a message to Kafka
func (m *MockProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	m.messages = append(m.messages, msg)
	return 0, 0, nil
}

// Close mocks closing the producer
func (m *MockProducer) Close() error {
	return nil
}

func TestSendMessage(t *testing.T) {
	mockProducer := &MockProducer{}

	message := &sarama.ProducerMessage{
		Topic: "meu-topico",
		Value: sarama.StringEncoder("Huncoding"),
	}

	partition, offset, err := mockProducer.SendMessage(message)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), partition)
	assert.Equal(t, int64(0), offset)
	assert.Equal(t, 1, len(mockProducer.messages))
	assert.Equal(t, "Huncoding", string(mockProducer.messages[0].Value.(sarama.StringEncoder)))
}
