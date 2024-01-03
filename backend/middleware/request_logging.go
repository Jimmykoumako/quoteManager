// middleware/request_logging.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// LogRequest logs information about incoming requests
func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Received %s request for %s from %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
		)

		// Continue to the next middleware or request handler
		c.Next()
	}
}
