package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "my-topic",
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("hello world"),
	})
	if err != nil {
		log.Fatal("cannot write message: ", err)
	}
}