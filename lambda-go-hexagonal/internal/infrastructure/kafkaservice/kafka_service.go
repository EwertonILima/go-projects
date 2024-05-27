package kafkaservice

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaAdapter interface for producing messages to Kafka
type KafkaService interface {
	Produce(msg []byte, topic string) error
}

// KafkaProducer manages the connection and message production to Kafka
type KafkaServiceImpl struct {
	Producer *kafka.Producer
}

// NewKafkaProducer creates a new KafkaProducer instance
// func NewKafkaProducer(config *kafka.ConfigMap) (*KafkaServiceImpl, error) {
// 	producer, err := kafka.NewProducer(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &KafkaServiceImpl{Producer: producer}, nil
// }

// // Close closes the Kafka producer connection
// func (p *KafkaServiceImpl) Close() {
// 	p.Producer.Close()
// }

// // Produce sends a message to Kafka
// func (p *KafkaServiceImpl) Produce(msg []byte, topic string) error {
// 	message := &kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          msg,
// 	}
// 	return p.Producer.Produce(message, nil)
// }


func (p *KafkaServiceImpl) Produce(msg []byte, topic string) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
	kafkaProducer, err := kafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	defer kafkaProducer.Close()
	kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)
}