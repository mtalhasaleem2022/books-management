package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"books-management-go/internal/models"
	"books-management-go/internal/repositories"
)

type BookService interface {
	GetBooks(ctx context.Context, limit, offset int) ([]models.Book, error)
	GetBookByID(ctx context.Context, id uint) (models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) error
	UpdateBook(ctx context.Context, id uint, book *models.Book) error
	DeleteBook(ctx context.Context, id uint) error
}

type bookService struct {
	repo     repositories.BookRepository
	cache    *repositories.RedisClient
	producer *repositories.KafkaProducer
}

func NewBookService(repo repositories.BookRepository, cache *repositories.RedisClient, producer *repositories.KafkaProducer) BookService {
	return &bookService{repo, cache, producer}
}

func (s *bookService) GetBooks(ctx context.Context, limit, offset int) ([]models.Book, error) {
	cacheKey := fmt.Sprintf("books:%d:%d", limit, offset)
	cached, err := s.cache.Get(ctx, cacheKey)
	if err == nil {
		var books []models.Book
		if json.Unmarshal([]byte(cached), &books) == nil {
			return books, nil
		}
	}

	books, err := s.repo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	serialized, _ := json.Marshal(books)
	s.cache.Set(ctx, cacheKey, string(serialized), 10*time.Minute)
	return books, nil
}

func (s *bookService) GetBookByID(ctx context.Context, id uint) (models.Book, error) {
	cacheKey := fmt.Sprintf("book:%d", id)
	cached, err := s.cache.Get(ctx, cacheKey)
	if err == nil {
		var book models.Book
		if json.Unmarshal([]byte(cached), &book) == nil {
			return book, nil
		}
	}

	book, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return models.Book{}, err
	}

	serialized, _ := json.Marshal(book)
	s.cache.Set(ctx, cacheKey, string(serialized), 5*time.Minute)
	return book, nil
}

func (s *bookService) CreateBook(ctx context.Context, book *models.Book) error {
	if err := book.Validate(); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, book); err != nil {
		return err
	}

	event := repositories.BookEvent{
		Type:      "BOOK_CREATED",
		Book:      *book,
		Timestamp: time.Now().Unix(),
	}
	s.producer.Publish(event)

	s.cache.DeleteByPattern(ctx, "books:*")
	return nil
}

func (s *bookService) UpdateBook(ctx context.Context, id uint, book *models.Book) error {
	if err := book.Validate(); err != nil {
		return err
	}

	if err := s.repo.Update(ctx, id, book); err != nil {
		return err
	}

	event := repositories.BookEvent{
		Type:      "BOOK_UPDATED",
		Book:      *book,
		Timestamp: time.Now().Unix(),
	}
	s.producer.Publish(event)

	s.cache.Delete(ctx, fmt.Sprintf("book:%d", id))
	s.cache.DeleteByPattern(ctx, "books:*")
	return nil
}

func (s *bookService) DeleteBook(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	event := repositories.BookEvent{
		Type:      "BOOK_DELETED",
		BookID:    id,
		Timestamp: time.Now().Unix(),
	}
	s.producer.Publish(event)

	s.cache.Delete(ctx, fmt.Sprintf("book:%d", id))
	s.cache.DeleteByPattern(ctx, "books:*")
	return nil
}
