// Package repository defines interfaces for interacting with various data entities in the Gtracker application.
package repository

import (
	"context"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/model"
)

// Goal represents the repository interface for managing goals in the database.
type Goal interface {
	// Create creates a new goal associated with a user in the database.
	Create(ctx context.Context, userId int, goal dto.CreateGoal) (int, error)

	// Get retrieves a list of goals associated with a user from the database.
	Get(ctx context.Context, userId int, listParams dto.ListParams) ([]model.Goal, error)

	// GetByID retrieves a goal by its ID from the database.
	GetByID(ctx context.Context, goalId int) (model.Goal, error)

	// UpdateByID updates a goal by its ID in the database.
	UpdateByID(ctx context.Context, goalId int, update dto.UpdateGoal) error

	// DeleteByID deletes a goal by its ID from the database.
	DeleteByID(ctx context.Context, goalId int) error
}

// Task represents the repository interface for managing tasks in the database.
type Task interface {
	// Create creates a new task associated with a goal in the database.
	Create(ctx context.Context, goalId int, task dto.CreateTask) (int, error)

	// Get retrieves a list of tasks associated with a goal from the database.
	Get(ctx context.Context, goalId int, listParams dto.ListParams) ([]model.Task, error)

	// GetByID retrieves a task by its ID from the database.
	GetByID(ctx context.Context, taskId int) (model.Task, error)

	// UpdateByID updates a task by its ID in the database.
	UpdateByID(ctx context.Context, taskId int, task dto.UpdateTask) error

	// DeleteByID deletes a task by its ID from the database.
	DeleteByID(ctx context.Context, taskId int) error
}

// User represents the repository interface for managing users in the database.
type User interface {
	// SignUp creates a new user account in the database.
	SignUp(ctx context.Context, up dto.SignUp) error

	// GetUserByName retrieves a user by their username from the database.
	GetUserByName(ctx context.Context, username string) (model.User, error)
}
