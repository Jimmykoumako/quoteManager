// utils/jwt_utils.go
package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// AccessTokenExpiration defines the expiration duration for access tokens
const AccessTokenExpiration = time.Hour

// RefreshTokenSecret is the secret key used to sign refresh tokens
var RefreshTokenSecret = []byte("your_refresh_token_secret_key")

// AccessTokenSecret is the secret key used to sign access tokens
var AccessTokenSecret = []byte("your_access_token_secret_key")

// ValidateRefreshToken validates the provided refresh token
func ValidateRefreshToken(refreshToken string) (bool, string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return RefreshTokenSecret, nil
	})

	if err != nil {
		return false, "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return false, "", errors.New("invalid token")
	}

	return true, claims.Subject, nil
}

// GenerateJWT generates a new JWT with the specified user ID as the subject
func GenerateJWTwithID(userID string) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(AccessTokenExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(AccessTokenSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
