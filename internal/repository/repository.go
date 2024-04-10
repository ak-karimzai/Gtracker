package repository

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository/goal"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository/task"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository/user"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/db"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
)

type Repository struct {
	Goal
	Task
	User
}

func NewRepository(db *db.DB, logger logger.Logger) *Repository {
	return &Repository{
		Goal: goal.NewRepository(db, logger),
		Task: task.NewRepository(db, logger),
		User: user.NewRepository(db, logger),
	}
}
