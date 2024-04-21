// Package auth_token provides functionality for creating and verifying JWT authentication tokens.
package auth_token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	MinSecretKeyLen = 12 // MinSecretKeyLen defines the minimum length for the secret key.
)

// JWTToken represents a JWT authentication token generator.
type JWTToken struct {
	secretKey      string        // Secret key used for signing the token.
	validationTime time.Duration // Duration for token validation.
}

// NewJWTToken creates a new JWTToken instance with the provided secret key and validation duration.
// It returns an error if the secret key length is less than the minimum required length.
func NewJWTToken(secretKey string, duration time.Duration) (*JWTToken, error) {
	if len(secretKey) < MinSecretKeyLen {
		return nil, fmt.Errorf(
			"invalid secret key: {%s} min len must be: %d",
			secretKey, MinSecretKeyLen)
	}
	return &JWTToken{secretKey: secretKey, validationTime: duration}, nil
}

// CreateToken creates a new JWT token with the provided user ID and username.
// It returns the JWT token as a string and an error if token creation fails.
func (tkn *JWTToken) CreateToken(userId int, username string) (string, error) {
	payload := &Payload{
		UserID:    userId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(tkn.validationTime),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(tkn.secretKey))
	return token, err
}

// VerifyToken verifies the authenticity of the provided JWT token.
// It returns the token payload if valid, otherwise returns an error.
func (t *JWTToken) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(tkn *jwt.Token) (any, error) {
		_, ok := tkn.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(t.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
