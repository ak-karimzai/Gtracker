// Package service provides the main service for managing goals, tasks, and authentication.
package service

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/auth"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/goal"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/task"
	auth_token "git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/auth-token"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
)

// Service provides the main service for managing goals, tasks, and authentication.
type Service struct {
	Goal             // Interface for managing goals.
	Task             // Interface for managing tasks.
	Auth             // Interface for authentication.
	auth_token.Maker // Interface for creating authentication tokens.
}

// NewService creates a new instance of the main service.
func NewService(repos *repository.Repository, tokenMaker auth_token.Maker, logger logger.Logger) *Service {
	return &Service{
		Goal:  goal.NewService(repos, logger),             // Initialize the goal service.
		Task:  task.NewService(repos, logger),             // Initialize the task service.
		Auth:  auth.NewService(repos, tokenMaker, logger), // Initialize the authentication service.
		Maker: tokenMaker,                                 // Set the authentication token maker.
	}
}
