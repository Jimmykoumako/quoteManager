// middleware/authorize.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"api/models"
)

// Authorize ensures that the authenticated user has the necessary permissions
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Fetch the user role from the authentication context
		userRole := getUserRoleFromContext(c)

		// Check if the user has the necessary permissions (adjust based on your authorization logic)
		if userRole != models.AdminRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		// Continue to the next middleware or request handler
		c.Next()
	}
}
