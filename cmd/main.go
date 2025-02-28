package main

import (
	"books-management-go/internal/services"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"books-management-go/config"
	"books-management-go/internal/controllers"
	"books-management-go/internal/repositories"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Books API
// @version 1.0
// @description REST API for managing books
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := config.LoadConfig()

	// Initialize dependencies
	db := repositories.InitDB(cfg.DBUrl)
	redisClient := repositories.NewRedisClient(cfg.RedisUrl)
	kafkaProducer := repositories.NewKafkaProducer(cfg.KafkaBrokers)

	// Repository
	bookRepo := repositories.NewBookRepository(db)

	// Service
	bookService := services.NewBookService(bookRepo, redisClient, kafkaProducer)

	// Controller
	bookController := controllers.NewBookController(bookService)

	// Gin Server
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes
	api := router.Group("/api/v1")
	{
		api.GET("/books", bookController.GetBooks)
		api.GET("/books/:id", bookController.GetBookByID)
		api.POST("/books", bookController.CreateBook)
		api.PUT("/books/:id", bookController.UpdateBook)
		api.DELETE("/books/:id", bookController.DeleteBook)
	}

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Cleanup
	redisClient.Close()
	kafkaProducer.Close()
	log.Println("Server exiting")
}
