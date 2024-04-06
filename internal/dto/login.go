package dto

import "fmt"

const (
	MaxUsernameLen = 64
	MinUsernameLen = 6
	MaxPasswordLen = 64
	MinPasswordLen = 6
)

// Login
// @Description Login request credentials
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login Login) Validate() error {
	if length := len(login.Username); length < MinUsernameLen ||
		length > MaxUsernameLen {
		return fmt.Errorf(
			"incorrect username {%s} length %d", login.Username, length)
	}

	if length := len(login.Password); length < MinPasswordLen ||
		length > MaxPasswordLen {
		return fmt.Errorf(
			"incorrect username {%s} length %d", login.Username, length)
	}
	return nil
}
