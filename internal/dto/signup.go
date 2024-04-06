package dto

import "fmt"

// SignUp
// @Description SignUp request credentials
type SignUp struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (signup SignUp) Validate() error {
	if length := len(signup.FirstName); length > MaxNameLen ||
		length < MinNameLen {
		return fmt.Errorf(
			"incorrect first name {%s} length %d", signup.FirstName, length)
	}

	if length := len(signup.LastName); length > MaxNameLen ||
		length < MinNameLen {
		return fmt.Errorf(
			"incorrect last name {%s} length %d", signup.LastName, length)
	}

	if length := len(signup.Username); length < MinUsernameLen ||
		length > MaxUsernameLen {
		return fmt.Errorf(
			"incorrect username {%s} length %d", signup.LastName, length)
	}

	if length := len(signup.Password); length < MinPasswordLen ||
		length > MaxPasswordLen {
		return fmt.Errorf(
			"incorrect password {%s} length %d", signup.LastName, length)
	}
	return nil
}
