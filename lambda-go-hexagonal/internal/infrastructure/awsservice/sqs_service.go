package awsservice

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
)

type SqsService interface {
	HandleAwsEvent(event events.SQSEvent) (string)
}

type SqsServiceImpl struct{}

func (s *SqsServiceImpl) HandleSQSEvent(event events.SQSEvent) (string) {
	var message string
	for _, record := range event.Records {
		log.Printf("Extracting message: %s", record.Body)
		message = record.Body
	}
	return message
}
