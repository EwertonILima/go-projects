package main

import (
	"context"
	"testing"

	"github.com/segmentio/kafka-go"
)

type mockKafkaWriter struct {
	messages []kafka.Message
	err      error
}

func (m *mockKafkaWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	m.messages = append(m.messages, msgs...)
	return m.err
}

func TestKafkaWriter(t *testing.T) {
	mockWriter := &mockKafkaWriter{}

	// Replace the actual writer with the mock writer
	writer := mockWriter

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("hello world"),
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(mockWriter.messages) != 1 {
		t.Errorf("expected 1 message, got %d", len(mockWriter.messages))
	}

	if string(mockWriter.messages[0].Value) != "hello world" {
		t.Errorf("expected message to be 'hello world', got %s", string(mockWriter.messages[0].Value))
	}
}
