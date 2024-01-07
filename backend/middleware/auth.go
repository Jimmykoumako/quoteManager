// middleware/authentication.go
package middleware

import (
	"api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware is a middleware to authenticate users using JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil || !utils.VerifyJWT(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
