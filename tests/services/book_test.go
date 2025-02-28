package services_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"books-management-go/config"
	"books-management-go/internal/models"
	"books-management-go/internal/repositories"
	"books-management-go/internal/repositories/mocks"
	"books-management-go/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestBookService_GetBooks(t *testing.T) {
	// Create a mock BookRepository
	mockRepo := &mocks.MockBookRepository{
		FindAllFunc: func(ctx context.Context, limit, offset int) ([]models.Book, error) {
			return []models.Book{
				{ID: 1, Title: "Test Book 1"},
				{ID: 2, Title: "Test Book 2"},
			}, nil
		},
	}

	// Create a mock RedisClient
	mockCache := &mocks.MockRedisClient{
		GetFunc: func(ctx context.Context, key string) (string, error) {
			return "", nil // Simulate cache miss
		},
		SetFunc: func(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
			return nil // Simulate successful cache set
		},
	}

	// Create a mock KafkaProducer
	mockKafka := &mocks.MockKafkaProducer{
		PublishFunc: func(event interface{}) {
			// Simulate successful Kafka event publishing
		},
	}

	fmt.Println(mockRepo, mockKafka, mockCache)

	cfg := config.LoadConfig()

	redisClient := repositories.NewRedisClient(cfg.RedisUrl)
	kafkaProducer := repositories.NewKafkaProducer(cfg.KafkaBrokers)

	// Create the service with the mock dependencies
	service := services.NewBookService(mockRepo, redisClient, kafkaProducer)

	// Test the GetBooks method
	books, err := service.GetBooks(context.Background(), 10, 0)
	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.Equal(t, "Test Book 1", books[0].Title)
	assert.Equal(t, "Test Book 2", books[1].Title)
}
