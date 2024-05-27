package main

import (
	"interface-kafka/internal/infrastructure/kafkaservice"
	"log"
	"fmt"
	"interface-kafka/internal/config"


)

func main() {
	// kafkaAdapter := kafkaservice.KafkaServiceImpl{}
	// kafkaAdapter.Produce([]byte("Test interface"), "my-topic")


	kafkaConfig, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    // Create Kafka producer instance using the loaded configuration
    kafkaProducer, err := kafkaservice.NewKafkaProducer(kafkaConfig)
    if err != nil {
        log.Fatal("Error creating Kafka producer:", err)
    }

    // Use the producer to send message
    err = kafkaProducer.Produce([]byte("Test interface"), "my-topic")
    if err != nil {
        log.Printf("Error producing message to Kafka: %v", err)
    } else {
        fmt.Println("Message sent to Kafka")
    }
}