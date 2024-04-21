// Package handler defines interfaces for handling HTTP requests related to goals, tasks, and authentication in the Gtracker application.
package handler

import "github.com/gin-gonic/gin"

// Goal interface defines methods for handling HTTP requests related to goals.
type Goal interface {
	// Create handles HTTP POST requests to create a new goal.
	Create(ctx *gin.Context)
	// Get handles HTTP GET requests to retrieve all goals.
	Get(ctx *gin.Context)
	// GetByID handles HTTP GET requests to retrieve a specific goal by its ID.
	GetByID(ctx *gin.Context)
	// UpdateByID handles HTTP PATCH requests to update a specific goal by its ID.
	UpdateByID(ctx *gin.Context)
	// DeleteByID handles HTTP DELETE requests to delete a specific goal by its ID.
	DeleteByID(ctx *gin.Context)
}

// Task interface defines methods for handling HTTP requests related to tasks.
type Task interface {
	// Create handles HTTP POST requests to create a new task.
	Create(ctx *gin.Context)
	// Get handles HTTP GET requests to retrieve all tasks associated with a goal.
	Get(ctx *gin.Context)
	// GetByID handles HTTP GET requests to retrieve a specific task by its ID.
	GetByID(ctx *gin.Context)
	// UpdateByID handles HTTP PUT requests to update a specific task by its ID.
	UpdateByID(ctx *gin.Context)
	// DeleteByID handles HTTP DELETE requests to delete a specific task by its ID.
	DeleteByID(ctx *gin.Context)
}

// Auth interface defines methods for handling HTTP requests related to user authentication.
type Auth interface {
	// SignUp handles HTTP POST requests for user registration.
	SignUp(ctx *gin.Context)
	// Login handles HTTP POST requests for user login and authentication.
	Login(ctx *gin.Context)
}
