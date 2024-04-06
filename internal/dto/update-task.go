package dto

import "fmt"

// UpdateTask
// @Description Update task request credentials
type UpdateTask struct {
	Name        *string    `json:"name" minLength:"3" maxLength:"64" example:"Run"`
	Description *string    `json:"description" minLength:"0" maxLength:"256" example:"Running everyday early morning"`
	Frequency   *Frequency `json:"frequency" example:"Daily"`
}

func (task UpdateTask) Validate() error {
	changes := 0
	if task.Name != nil {
		if length := len(*task.Name); length > MaxNameLen ||
			length < MinNameLen {
			return fmt.Errorf(
				"incorrect name {%s} length %d", *task.Name, length)
		}
		changes++
	}

	if task.Description != nil {
		if length := len(*task.Description); length > MaxDescriptionLen ||
			length < MinDescriptionLen {
			return fmt.Errorf(
				"incorrect description {%s} length %d", *task.Description, length)
		}
		changes++
	}

	if task.Frequency != nil {
		if !task.Frequency.Validate() {
			return fmt.Errorf(
				"incorrect completion status {%s}", *task.Frequency)
		}
		changes++
	}

	if changes == 0 {
		return ErrEmptyUpdate
	}

	return nil
}
