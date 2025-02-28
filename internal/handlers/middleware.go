package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs the details of each request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process the request
		c.Next()

		// Log the request details
		log.Printf(
			"[%s] %s %s %s %d %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Request.Proto,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

// ErrorHandlingMiddleware handles errors and returns a consistent error response
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next()

		// Handle errors from the request
		if len(c.Errors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": c.Errors.Last().Error(),
			})
		}
	}
}

// RecoveryMiddleware recovers from panics and prevents the server from crashing
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("RecoveryMiddleware: recovered from panic: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}
