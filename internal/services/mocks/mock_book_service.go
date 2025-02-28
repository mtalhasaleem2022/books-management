package mocks

import (
	"context"
	"books-management-go/internal/models"
)

// MockBookService is a mock implementation of the BookService interface
type MockBookService struct {
	GetBooksFunc    func(ctx context.Context, limit, offset int) ([]models.Book, error)
	GetBookByIDFunc func(ctx context.Context, id uint) (models.Book, error)
	CreateBookFunc  func(ctx context.Context, book *models.Book) error
	UpdateBookFunc  func(ctx context.Context, id uint, book *models.Book) error
	DeleteBookFunc  func(ctx context.Context, id uint) error
}

// GetBooks calls the mocked GetBooksFunc
func (m *MockBookService) GetBooks(ctx context.Context, limit, offset int) ([]models.Book, error) {
	return m.GetBooksFunc(ctx, limit, offset)
}

// GetBookByID calls the mocked GetBookByIDFunc
func (m *MockBookService) GetBookByID(ctx context.Context, id uint) (models.Book, error) {
	return m.GetBookByIDFunc(ctx, id)
}

// CreateBook calls the mocked CreateBookFunc
func (m *MockBookService) CreateBook(ctx context.Context, book *models.Book) error {
	return m.CreateBookFunc(ctx, book)
}

// UpdateBook calls the mocked UpdateBookFunc
func (m *MockBookService) UpdateBook(ctx context.Context, id uint, book *models.Book) error {
	return m.UpdateBookFunc(ctx, id, book)
}

// DeleteBook calls the mocked DeleteBookFunc
func (m *MockBookService) DeleteBook(ctx context.Context, id uint) error {
	return m.DeleteBookFunc(ctx, id)
}