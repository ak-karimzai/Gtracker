// Package service provides interfaces for managing goals, tasks, and authentication.
package service

import (
	"context"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/model"
)

// Goal defines methods for managing goals.
type Goal interface {
	// Create creates a new goal for a user.
	Create(ctx context.Context, userId int, goal dto.CreateGoal) (model.Goal, error)

	// Get retrieves goals for a user based on provided parameters.
	Get(ctx context.Context, userId int, listParams dto.ListParams) ([]model.Goal, error)

	// GetByID retrieves a goal by its ID.
	GetByID(ctx context.Context, userId, goalId int) (model.Goal, error)

	// UpdateByID updates a goal by its ID.
	UpdateByID(ctx context.Context, userId, goalId int, update dto.UpdateGoal) error

	// DeleteByID deletes a goal by its ID.
	DeleteByID(ctx context.Context, userId, goalId int) error
}

// Task defines methods for managing tasks.
type Task interface {
	// Create creates a new task for a goal.
	Create(ctx context.Context, userId, goalId int, task dto.CreateTask) (model.Task, error)

	// Get retrieves tasks for a goal based on provided parameters.
	Get(ctx context.Context, userId, goalId int, listParams dto.ListParams) ([]model.Task, error)

	// GetByID retrieves a task by its ID.
	GetByID(ctx context.Context, userId, goalId, taskId int) (model.Task, error)

	// UpdateByID updates a task by its ID.
	UpdateByID(ctx context.Context, userId, goalId, taskId int, task dto.UpdateTask) error

	// DeleteByID deletes a task by its ID.
	DeleteByID(ctx context.Context, userId, goalId, taskId int) error
}

// Auth defines methods for authentication.
type Auth interface {
	// SignUp registers a new user.
	SignUp(ctx context.Context, up dto.SignUp) error

	// Login authenticates a user.
	Login(ctx context.Context, in dto.Login) (*dto.LoginResponse, error)
}
