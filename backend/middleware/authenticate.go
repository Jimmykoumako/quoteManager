// File: middleware/authenticate.go

package middleware

import (
	"github.com/gin-gonic/gin"
)

// SetUsernameToContext sets the username in the Gin context
func SetUsernameToContext(c *gin.Context, username string) {
	c.Set("username", username)
}

// SetUserRoleToContext sets the user role in the Gin context
func SetUserRoleToContext(c *gin.Context, userRole string) {
	c.Set("user_role", userRole)
}

// GetUsernameFromContext retrieves the username from the Gin context
func GetUsernameFromContext(c *gin.Context) string {
	// Assuming the username is stored in the Gin context under the key "username"
	username, exists := c.Get("username")
	if !exists {
		// Handle the case where the username is not found in the context
		return ""
	}

	// Convert the retrieved username to a string (assuming it's a string)
	if usernameStr, ok := username.(string); ok {
		return usernameStr
	}

	// Handle the case where the username is not of the expected type
	return ""
}

// GetUserRoleFromContext retrieves the user role from the Gin context
func GetUserRoleFromContext(c *gin.Context) string {
	// Assuming the user role is stored in the Gin context under the key "user_role"
	userRole, exists := c.Get("user_role")
	if !exists {
		// Handle the case where the user role is not found in the context
		return ""
	}

	// Convert the retrieved user role to a string (assuming it's a string)
	if userRoleStr, ok := userRole.(string); ok {
		return userRoleStr
	}

	// Handle the case where the user role is not of the expected type
	return ""
}
