package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"lambda-go-hexagonal/internal/domain/service"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.SQSEvent) {
	service := service.HandleEventsServiceImpl{}
	service.ProcessEvent(ctx, event)
}
