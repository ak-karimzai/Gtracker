package model

import "time"

// User represents a user entity in the Gtracker application.
type User struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password_hash"` // Password field is not included in JSON serialization.
	CreatedAt time.Time `json:"created_at" db:"create_at"`
}
