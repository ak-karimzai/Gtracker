package dto

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyUpdate = errors.New("nothing to change")
)

// UpdateGoal
// @Description Update goal request credentials
type UpdateGoal struct {
	Name             *string           `json:"name" minLength:"3" maxLength:"64" example:"Sport"`
	Description      *string           `json:"description" minLength:"3" maxLength:"64" example:"Workout everyday for healthy life"`
	CompletionStatus *CompletionStatus `json:"completion_status" example:"Skipped"`
	StartDate        *DDMMYYYY         `json:"start_date" example:"01-01-2023"`
	TargetDate       *DDMMYYYY         `json:"target_date" example:"01-01-2024"`
}

func (goal UpdateGoal) Validate() error {
	changes := 0
	if goal.Name != nil {
		if length := len(*goal.Name); length > MaxNameLen ||
			length < MinNameLen {
			return fmt.Errorf(
				"incorrect name {%s} length %d", *goal.Name, length)
		}
		changes++
	}

	if goal.Description != nil {
		if length := len(*goal.Description); length > MaxDescriptionLen ||
			length < MinDescriptionLen {
			return fmt.Errorf(
				"incorrect description {%s} length %d", *goal.Description, length)
		}
		changes++
	}

	if goal.CompletionStatus != nil {
		if !goal.CompletionStatus.Validate() {
			return fmt.Errorf(
				"incorrect completion status {%s}", *goal.CompletionStatus)
		}
		changes++
	}

	if goal.StartDate != nil {
		if !goal.StartDate.Validate() {
			return fmt.Errorf(
				"incorrect start date: %s", *goal.StartDate)
		}
		changes++
	}

	if goal.TargetDate != nil {
		if !goal.TargetDate.Validate() {
			return fmt.Errorf(
				"incorrect target date: %s", *goal.TargetDate)
		}
		changes++
	}

	if changes == 0 {
		return ErrEmptyUpdate
	}

	return nil
}
