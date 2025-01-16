package handlers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("[MIDDLEWARE] %s %s %d %s", c.Request.Method, c.Request.URL.String(), c.Writer.Status(), duration)
	}
}
