package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
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

var (
	// SecretKey is the secret key used for signing and verifying JWT tokens
	SecretKey = []byte("your_secret_key_here")
)

// Claims represents the claims of a JWT token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateAccessToken generates a new access token
func GenerateAccessToken(username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute) // Adjust the expiration time as needed

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateAccessToken validates an access token
func ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

