// File: controllers/auth_controller.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login simulates the login process
func Login(c *gin.Context) {
	// Assuming user authentication is successful, you don't need to explicitly trigger the middleware
	// If authentication fails, Gin's AuthMiddleware will handle it automatically

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
