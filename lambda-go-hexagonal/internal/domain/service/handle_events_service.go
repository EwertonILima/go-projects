package service

import (
	"context"
	"fmt"
	"log"

	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"lambda-go-hexagonal/internal/domain/entities"
	"lambda-go-hexagonal/internal/domain/mapper"
	"lambda-go-hexagonal/internal/infrastructure/awsservice"
	"lambda-go-hexagonal/internal/infrastructure/kafkaservice"
)

type HandleEventsService interface {
	ProcessEvent(ctx context.Context, event events.SQSEvent) error
}

type HandleEventsServiceImpl struct {
	sqsservice awsservice.SqsService
	kafka      kafkaservice.KafkaService
}
func NewHandleEventsService(sqsService awsservice.SqsService, kafkaService kafkaservice.KafkaService) *HandleEventsServiceImpl {
	return &HandleEventsServiceImpl{
		sqsservice: sqsService,
		kafka:      kafkaService,
	}
}

func (h *HandleEventsServiceImpl) ProcessEvent(ctx context.Context, event events.SQSEvent) error {
	// ... (same logic as before)
	var sqsEntity entities.SqsEntity

	sqsservice := awsservice.SqsServiceImpl{}
	sqsmessage := sqsservice.HandleSQSEvent(event)

	errjson := json.Unmarshal([]byte(sqsmessage), &sqsEntity)
	if errjson != nil {
		fmt.Println("Error unmarshalling JSON:", errjson)
	}

	kafkaMsg := mapper.MapSqsToOpenFinanceJson(sqsEntity)
	kafkaMsgBytes, _ := json.Marshal(kafkaMsg)

	log.Printf("Message to kafka: %s", kafkaMsgBytes)

	// Use the injected Kafka adapter to produce message
	kafkaAdapter := kafkaservice.KafkaServiceImpl{}
	err := kafkaAdapter.Produce(kafkaMsgBytes, "your-topic")
	if err != nil {
		log.Printf("Error producing message to Kafka: %v", err)
		return err
	}
	

	fmt.Println("Message sent to Kafka")
	return nil
}
