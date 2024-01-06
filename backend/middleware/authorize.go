// middleware/authorize.go

package middleware

import (
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getUserRoleFromContext retrieves the user role from the Gin context
func getUserRoleFromContext(c *gin.Context) string {
	// Implement your logic to get the user role from the context
	// For example, assuming the role is stored under the key "user_role"
	userRole, exists := c.Get("user_role")
	if !exists {
		return ""
	}

	// Convert to string if it's of the expected type
	if userRoleStr, ok := userRole.(string); ok {
		return userRoleStr
	}

	return ""
}

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

func isAuthenticated(c *gin.Context) bool {
	fmt.Printf("Welcome to isAuthenticated")
	// Check if the user is authenticated (e.g., verify token or session)
	// Return true if authenticated, false otherwise
	// Implement your own authentication logic
	return true
}
