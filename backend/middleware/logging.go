// middleware/logging.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// LogFeedbackAction logs information about feedback-related actions
func LogFeedbackAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log relevant information (user, action type, etc.)
		log.Printf("User %s performed %s action on feedback ID %s",
			getUsernameFromContext(c),
			c.Request.Method,
			c.Param("id"),
		)

		// Continue to the next middleware or request handler
		c.Next()
	}
}
