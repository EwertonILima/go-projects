package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event json.RawMessage) error {
    var eventAWS events.SQSEvent
    var s3Event events.S3Event

	log.Printf("Received message: %s", event)

    if err := json.Unmarshal(event, &eventAWS); err == nil {
		for _, record := range eventAWS.Records {
			if record.EventSource == "aws:sqs" {
				return handleSQSEvent(eventAWS)
			}
			if record.EventSource == "aws:s3" {
				
				if err := json.Unmarshal(event, &s3Event); err == nil {
					// Handle S3 event
					return handleS3Event(s3Event)
				}
			}
		}
        
    }

    // // Try to parse the event as an S3 event
    // if err := json.Unmarshal(event, &s3Event); err == nil {
    //     // Handle S3 event
    //     return handleS3Event(s3Event)
    // }

    return fmt.Errorf("unknown event type")
}

func handleSQSEvent(event events.SQSEvent) error {
    for _, message := range event.Records {
        log.Printf("Received message: %s", message.Body)
        // Process the message
    }
    return nil
}

func handleS3Event(event events.S3Event) error {
    for _, record := range event.Records {
        s3 := record.S3
        log.Printf("Bucket = %s, Key = %s", s3.Bucket.Name, s3.Object.Key)
        // Process the S3 event
    }
    return nil
}

func main() {
    lambda.Start(handler)
}