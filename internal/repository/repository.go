package repository

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository/goal"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository/task"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository/user"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/db"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
)

// Repository aggregates instances of goal, task, and user repositories.
type Repository struct {
	Goal
	Task
	User
}

// NewRepository creates a new Repository instance with the provided database connection and logger.
func NewRepository(db *db.DB, logger logger.Logger) *Repository {
	return &Repository{
		Goal: goal.NewRepository(db, logger),
		Task: task.NewRepository(db, logger),
		User: user.NewRepository(db, logger),
	}
}
