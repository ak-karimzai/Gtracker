package dto

import (
	"fmt"
)

const (
	MinNameLen        = 3
	MaxNameLen        = 64
	MinDescriptionLen = 0
	MaxDescriptionLen = 256
)

// CreateGoal goal create request
// @Description Create request requirments
type CreateGoal struct {
	Name        string   `json:"name"    binding:"required" minLength:"3" maxLength:"64" example:"Sport"`
	Description string   `json:"description" binding:"required" minLength:"0" maxLength:"256" example:"Workout everyday for healthy life"`
	StartDate   DDMMYYYY `json:"start_date"  binding:"required" example:"01-01-2023"`
	TargetDate  DDMMYYYY `json:"target_date" binding:"required" example:"01-01-2024"`
}

func (cg CreateGoal) Validate() error {
	if length := len(cg.Name); length > MaxNameLen ||
		length < MinNameLen {
		return fmt.Errorf(
			"incorrect username {%s} length %d", cg.Name, length)
	}

	if length := len(cg.Description); length < MinDescriptionLen ||
		length > MaxDescriptionLen {
		return fmt.Errorf(
			"incorrect username {%s} length %d", cg.Description, length)
	}

	if !cg.StartDate.Validate() {
		return fmt.Errorf(
			"incorrect start date: %s", cg.StartDate)
	}

	if !cg.TargetDate.Validate() {
		return fmt.Errorf(
			"incorrect target date: %s", cg.TargetDate)
	}

	return nil
}
