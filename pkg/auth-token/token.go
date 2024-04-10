package auth_token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	MinSecretKeyLen = 12
)

type JWTToken struct {
	secretKey      string
	validationTime time.Duration
}

func NewJWTToken(secretKey string, duration time.Duration) (*JWTToken, error) {
	if len(secretKey) < MinSecretKeyLen {
		return nil, fmt.Errorf(
			"invalid secret key: {%s} min len must be: %d",
			secretKey, MinSecretKeyLen)
	}
	return &JWTToken{secretKey: secretKey, validationTime: duration}, nil
}

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
