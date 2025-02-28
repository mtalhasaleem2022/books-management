package repositories

import (
	"books-management-go/internal/models"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// BookEvent represents an event related to a book
type BookEvent struct {
	Type      string      `json:"type"`      // Event type: "BOOK_CREATED", "BOOK_UPDATED", "BOOK_DELETED"
	Book      models.Book `json:"book"`      // Book data (for create/update events)
	BookID    uint        `json:"bookId"`    // Book ID (for delete events)
	Timestamp int64       `json:"timestamp"` // Event timestamp
}

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(brokers),
			Topic: "book_events",
		},
	}
}

func (kp *KafkaProducer) Publish(event BookEvent) {
	eventJSON, _ := json.Marshal(event)
	err := kp.writer.WriteMessages(context.Background(), kafka.Message{
		Value: eventJSON,
	})
	if err != nil {
		log.Printf("Failed to publish Kafka event: %v", err)
	}
}

func (kp *KafkaProducer) Close() {
	kp.writer.Close()
}
