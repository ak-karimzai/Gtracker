package model

import (
	"time"
)

type Goal struct {
	ID               int       `json:"id" example:"1"`
	Name             string    `json:"name" minLength:"3" maxLength:"64" example:"Sport"`
	Description      string    `json:"description" minLength:"3" maxLength:"64" example:"Workout everyday for healthy life"`
	CompletionStatus string    `json:"completion_status" example:"Progress"`
	StartDate        time.Time `json:"start_date" example:"2023-01-01T15:04:05Z"`
	TargetDate       time.Time `json:"target_date" example:"2024-01-01T15:04:05Z"`
	CreatedAt        time.Time `json:"created_at" example:"2023-04-01T15:04:05Z"`
	UpdatedAt        time.Time `json:"updated_at" example:"2023-04-01T15:04:05Z"`
	UserID           int       `json:"-"`
}
