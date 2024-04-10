package auth_token

import (
	"fmt"
	"time"
)

var ErrExpiredToken = fmt.Errorf("token expired")
var ErrInvalidToken = fmt.Errorf("invalid token")

type Payload struct {
	UserID    int       `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
