package repositories

import (
	"context"

	"books-management-go/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]models.Book, error)
	FindByID(ctx context.Context, id uint) (models.Book, error)
	Create(ctx context.Context, book *models.Book) error
	Update(ctx context.Context, id uint, book *models.Book) error
	Delete(ctx context.Context, id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) FindAll(ctx context.Context, limit, offset int) ([]models.Book, error) {
	var books []models.Book
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&books)
	return books, result.Error
}

func (r *bookRepository) FindByID(ctx context.Context, id uint) (models.Book, error) {
	var book models.Book
	result := r.db.WithContext(ctx).First(&book, id)
	return book, result.Error
}

func (r *bookRepository) Create(ctx context.Context, book *models.Book) error {
	return r.db.WithContext(ctx).Create(book).Error
}

func (r *bookRepository) Update(ctx context.Context, id uint, book *models.Book) error {
	return r.db.WithContext(ctx).Model(&models.Book{}).Where("id = ?", id).Updates(book).Error
}

func (r *bookRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Book{}, id).Error
}
