// Package util provides utility functions for common tasks.
package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash for the given password.
// It returns the hashed password as a string and an error if any.
func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPwd), nil
}

// CheckPassword checks if the provided password matches the hashed password.
// It returns an error if the passwords don't match.
func CheckPwd(pwd, hashedPwd string) error {
	return bcrypt.
		CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}
