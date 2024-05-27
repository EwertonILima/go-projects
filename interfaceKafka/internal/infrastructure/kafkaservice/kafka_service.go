package kafkaservice

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"interface-kafka/internal/config"
)

// KafkaAdapter interface for producing messages to Kafka
type KafkaService interface {
	Produce(msg []byte, topic string) error
}

// KafkaProducer manages the connection and message production to Kafka
type KafkaLibImpl struct {
	Producer *kafka.Producer
}

// NewKafkaProducer creates a new KafkaProducer instance
func NewKafkaProducer(config *config.Config) (*KafkaLibImpl, error) {
	kafkaConfig := &kafka.ConfigMap{
        "bootstrap.servers": config.BootstrapServers,
        // Add other required or desired Kafka configuration options
    }
    producer, err := kafka.NewProducer(kafkaConfig)
    if err != nil {
        return nil, err
    }
	return &KafkaLibImpl{Producer: producer}, nil
}

// Produce sends a message to Kafka
func (p *KafkaLibImpl) Produce(msg []byte, topic string) error {
	
	defer p.Producer.Close()

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}
	return p.Producer.Produce(message, nil)
}

