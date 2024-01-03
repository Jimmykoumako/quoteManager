package database

import "errors"

var (
	// ErrInvalidPayload is returned when the request payload is invalid
	ErrInvalidPayload = errors.New("invalid request payload")
)
