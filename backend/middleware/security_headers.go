// middleware/security_headers.go

package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeadersMiddleware sets security headers in the response
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		// Add more security headers as needed

		c.Next()
	}
}
