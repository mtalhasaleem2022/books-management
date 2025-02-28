package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"books-management-go/internal/controllers"
	"books-management-go/internal/models"
	"books-management-go/internal/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBookController_CreateBook(t *testing.T) {
	// Create a mock BookService
	mockService := &mocks.MockBookService{
		CreateBookFunc: func(ctx context.Context, book *models.Book) error {
			return nil // Simulate successful creation
		},
	}

	// Create the controller with the mock service
	controller := controllers.NewBookController(mockService)

	// Set up the Gin router
	router := gin.Default()
	router.POST("/books", controller.CreateBook)

	// Create a test book
	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2023,
	}
	body, _ := json.Marshal(book)

	// Create a request
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, resp.Code)
}
