package model

import "time"

//Name        *string    `json:"name" minLength:"3" maxLength:"64" example:"Run"`
//Description *string    `json:"description" minLength:"0" maxLength:"256" example:"Running everyday early morning"`
//Frequency   *Frequency `json:"frequency" example:"Daily"`

type Task struct {
	ID          int       `json:"id" example:"1""`
	Name        string    `json:"name" minLength:"3" maxLength:"64" example:"Run"`
	Description string    `json:"description" minLength:"0" maxLength:"256" example:"Running everyday early morning"`
	Frequency   string    `json:"frequency" example:"Daily"`
	CreatedAt   time.Time `json:"created_at" example:"2023-04-01T15:04:05Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-04-01T15:04:05Z"`
	GoalID      int       `json:"-"`
}
