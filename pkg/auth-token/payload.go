// Package auth_token provides functionality for creating and verifying authentication tokens.
package auth_token

import (
	"fmt"
	"time"
)

// ErrExpiredToken is returned when a token is expired.
var ErrExpiredToken = fmt.Errorf("token expired")

// ErrInvalidToken is returned when a token is invalid.
var ErrInvalidToken = fmt.Errorf("invalid token")

// Payload represents the payload of an authentication token.
type Payload struct {
	UserID    int       `json:"id"`         // User ID associated with the token.
	Username  string    `json:"username"`   // Username associated with the token.
	IssuedAt  time.Time `json:"issued_at"`  // Time when the token was issued.
	ExpiredAt time.Time `json:"expired_at"` // Time when the token expires.
}

// Valid checks if the token is valid.
// It returns nil if the token is valid, otherwise returns an error.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
