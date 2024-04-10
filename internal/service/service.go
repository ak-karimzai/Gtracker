package service

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/auth"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/goal"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/task"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/auth-token"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
)

type Service struct {
	Goal
	Task
	Auth
	auth_token.Maker
}

func NewService(repos *repository.Repository, tokenMaker auth_token.Maker, logger logger.Logger) *Service {
	return &Service{
		Goal:  goal.NewService(repos, logger),
		Task:  task.NewService(repos, logger),
		Auth:  auth.NewService(repos, tokenMaker, logger),
		Maker: tokenMaker,
	}
}
