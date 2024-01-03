// middleware/logging.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// getUsernameFromContext retrieves the username from the Gin context
func getUsernameFromContext(c *gin.Context) string {
	// Assuming the username is stored in the Gin context under the key "username"
	username, exists := c.Get("username")
	if !exists {
		return ""
	}

	// Convert to string if it's of the expected type
	if usernameStr, ok := username.(string); ok {
		return usernameStr
	}

	return ""
}

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


