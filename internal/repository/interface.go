package repository

import (
	"context"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/model"
)

type Goal interface {
	Create(ctx context.Context, userId int, goal dto.CreateGoal) (int, error)
	Get(ctx context.Context, userId int, listParams dto.ListParams) ([]model.Goal, error)
	GetByID(ctx context.Context, goalId int) (model.Goal, error)
	UpdateByID(ctx context.Context, goalId int, update dto.UpdateGoal) error
	DeleteByID(ctx context.Context, goalId int) error
}

type Task interface {
	Create(ctx context.Context, goalId int, task dto.CreateTask) (int, error)
	Get(ctx context.Context, goalId int, listParams dto.ListParams) ([]model.Task, error)
	GetByID(ctx context.Context, taskId int) (model.Task, error)
	UpdateByID(ctx context.Context, taskId int, task dto.UpdateTask) error
	DeleteByID(ctx context.Context, taskId int) error
}

type User interface {
	SignUp(ctx context.Context, up dto.SignUp) error
	GetUserByName(ctx context.Context, username string) (model.User, error)
}
