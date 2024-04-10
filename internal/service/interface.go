package service

import (
	"context"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/model"
)

type Goal interface {
	Create(ctx context.Context, userId int, goal dto.CreateGoal) (model.Goal, error)
	Get(ctx context.Context, userId int, listParams dto.ListParams) ([]model.Goal, error)
	GetByID(ctx context.Context, userId, goalId int) (model.Goal, error)
	UpdateByID(ctx context.Context, userId, goalId int, update dto.UpdateGoal) error
	DeleteByID(ctx context.Context, userId, goalId int) error
}

type Task interface {
	Create(ctx context.Context, userId, goalId int, task dto.CreateTask) (model.Task, error)
	Get(ctx context.Context, userId, goalId int, listParams dto.ListParams) ([]model.Task, error)
	GetByID(ctx context.Context, userId, goalId, taskId int) (model.Task, error)
	UpdateByID(ctx context.Context, userId, goalId, taskId int, task dto.UpdateTask) error
	DeleteByID(ctx context.Context, userId, goalId, taskId int) error
}

type Auth interface {
	SignUp(ctx context.Context, up dto.SignUp) error
	Login(ctx context.Context, in dto.Login) (*dto.LoginResponse, error)
}
