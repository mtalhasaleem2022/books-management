package models

import "github.com/go-playground/validator/v10"

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title" validate:"required,min=3"`
	Author string `json:"author" validate:"required,min=3"`
	Year   int    `json:"year" validate:"required,gt=1900"`
}

func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
