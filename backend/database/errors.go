package database

import "errors"

var (
	// ErrInvalidPayload is returned when the request payload is invalid
	ErrInvalidPayload = errors.New("invalid request payload")
)

// In a custom errors package or database package
var ErrUsernameExists = errors.New("username already in use")