package mocks

import (
	"context"

	"books-management-go/internal/models"
)

// MockBookRepository is a mock implementation of the BookRepository interface
type MockBookRepository struct {
	FindAllFunc  func(ctx context.Context, limit, offset int) ([]models.Book, error)
	FindByIDFunc func(ctx context.Context, id uint) (models.Book, error)
	CreateFunc   func(ctx context.Context, book *models.Book) error
	UpdateFunc   func(ctx context.Context, id uint, book *models.Book) error
	DeleteFunc   func(ctx context.Context, id uint) error
}

func (m *MockBookRepository) FindAll(ctx context.Context, limit, offset int) ([]models.Book, error) {
	return m.FindAllFunc(ctx, limit, offset)
}

func (m *MockBookRepository) FindByID(ctx context.Context, id uint) (models.Book, error) {
	return m.FindByIDFunc(ctx, id)
}

func (m *MockBookRepository) Create(ctx context.Context, book *models.Book) error {
	return m.CreateFunc(ctx, book)
}

func (m *MockBookRepository) Update(ctx context.Context, id uint, book *models.Book) error {
	return m.UpdateFunc(ctx, id, book)
}

func (m *MockBookRepository) Delete(ctx context.Context, id uint) error {
	return m.DeleteFunc(ctx, id)
}
