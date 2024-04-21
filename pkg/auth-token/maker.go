// Package auth_token provides functionality for creating and verifying authentication tokens.
package auth_token

// Maker defines the interface for creating and verifying authentication tokens.
type Maker interface {
	// CreateToken generates a new authentication token for the given user ID and username.
	CreateToken(userId int, username string) (string, error)

	// VerifyToken verifies the authenticity of the provided token and returns the payload if valid.
	VerifyToken(token string) (*Payload, error)
}
