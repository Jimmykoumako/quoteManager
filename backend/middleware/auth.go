// middleware/auth.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the request is authenticated
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Check authentication logic here
        // For example, verify an authentication token or session

        // If authentication is successful, proceed to the next middleware or handler
        // Otherwise, return an unauthorized response
        if isAuthenticated(c) {
            c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
        }
    }
}


