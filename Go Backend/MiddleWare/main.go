package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware that uses a Goroutine to log asynchronously
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// process request
		c.Next()

		go func() {
			endTime := time.Now()
			latency := endTime.Sub(startTime)
			status := c.Writer.Status()
			fmt.Printf("Status: %d | Latency: %v | Path: %s\n", status, latency, c.Request.URL.Path)
		}()
	}
}

// Authentication Middleware that should be blocked until verified
func AuthMiddlewware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fix the typo in the header key
		token := c.GetHeader("Authorization")
		print("token = ", token)
		if token != "Bearer mysecretoken" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(LoggingMiddleware())
	r.Use(AuthMiddlewware())

	r.GET("/books", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "List of books"})
	})

	//  curl -H "Authorization: Bearer mysecretoken" http://localhost:5080/books

	r.Run(":5080")
}
