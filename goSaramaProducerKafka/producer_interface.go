// producer.go
package main

import "github.com/IBM/sarama"

// Producer is an interface for sending messages to Kafka
type Producer interface {
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
	Close() error
}

// SaramaProducer is a wrapper around the Sarama producer
type SaramaProducer struct {
	producer sarama.SyncProducer
}

// NewSaramaProducer creates a new SaramaProducer
func NewSaramaProducer(brokers []string, config *sarama.Config) (Producer, error) {
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &SaramaProducer{producer: producer}, nil
}

// SendMessage sends a message to Kafka
func (p *SaramaProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return p.producer.SendMessage(msg)
}

// Close closes the producer
func (p *SaramaProducer) Close() error {
	return p.producer.Close()
}
