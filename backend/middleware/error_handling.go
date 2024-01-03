// middleware/error_handling.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ErrorHandlerMiddleware handles errors and returns a standardized response
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Recovery] %s\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()

		c.Next()
	}
}
