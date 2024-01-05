// File: controllers/auth_controller.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login simulates the login process
func Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
