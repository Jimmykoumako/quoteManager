// middleware/db.go
package middleware

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "api/database"
)

// DatabaseMiddleware initializes the database connection and sets it in the context
func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        database.SetDB(db)
        defer database.CloseDB()
        c.Next()
    }
}
