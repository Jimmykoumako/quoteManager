package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserIDFromContext retrieves the user ID from the Gin context
func GetUserIDFromContext(c *gin.Context) string {
	// Assuming the user ID is stored in the Gin context under the key "UserID"
	userID, exists := c.Get("UserID")
	if !exists {
		// Handle the case where the user ID is not found in the context
		return ""
	}

	// Convert the retrieved user ID to a string (assuming it's a string)
	if userIDStr, ok := userID.(string); ok {
		return userIDStr
	}

	// Handle the case where the user ID is not of the expected type
	return ""
}

// ConvertUserIDToUint converts a user ID from string to uint
func ConvertUserIDToUint(userID string) (uint, error) {
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(userIDUint), nil
}
