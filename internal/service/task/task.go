package task

import (
	"context"
	"errors"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/model"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository"
	service_errors "git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/service-errors"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/db"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
)

type Service struct {
	repo   *repository.Repository
	logger logger.Logger
}

func NewService(repo *repository.Repository, logger logger.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s Service) Create(ctx context.Context, userId, goalId int, task dto.CreateTask) (model.Task, error) {
	if err := task.Validate(); err != nil {
		s.logger.Error(err)
		return model.Task{}, service_errors.ErrInvalidCredentials
	}

	goal, err := s.repo.Goal.GetByID(ctx, goalId)
	if err != nil {
		s.logger.Error(err)
		return model.Task{}, err
	}

	if goal.UserID != userId {
		s.logger.Printf("permission denied: %d != %d", goal.UserID, userId)
		return model.Task{}, service_errors.ErrPermissionDenied
	}

	id, err := s.repo.Task.Create(ctx, goalId, task)
	if err != nil {
		s.logger.Error(err)
		var finalErr = service_errors.ErrServiceNotAvailable
		if errors.Is(err, db.ErrConflict) {
			finalErr = service_errors.ErrAlreadyExists
		}
		return model.Task{}, finalErr
	}

	taskFromDB, err := s.repo.Task.GetByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		var finalErr = service_errors.ErrServiceNotAvailable
		if errors.Is(err, db.ErrNotFound) {
			finalErr = service_errors.ErrNotFound
		}
		return model.Task{}, finalErr
	}
	return taskFromDB, nil
}

func (s Service) Get(ctx context.Context, userId, goalId int, listParams dto.ListParams) ([]model.Task, error) {
	if err := listParams.Validate(); err != nil {
		s.logger.Error(err)
		return []model.Task{}, service_errors.ErrInvalidCredentials
	}

	goal, err := s.repo.Goal.GetByID(ctx, goalId)
	if err != nil {
		s.logger.Error(err)
		return []model.Task{}, err
	}

	if goal.UserID != userId {
		s.logger.Print("Error (Get): permission denied -> Goal user id are not equal")
		return []model.Task{}, service_errors.ErrPermissionDenied
	}

	tasks, err := s.repo.Task.Get(ctx, goalId, listParams)
	if err != nil {
		s.logger.Error(err)
		var finalErr = service_errors.ErrServiceNotAvailable
		if errors.Is(err, db.ErrNotFound) {
			finalErr = service_errors.ErrNotFound
		}
		return []model.Task{}, finalErr
	}
	return tasks, nil
}

func (s Service) GetByID(ctx context.Context, userId, goalId, taskId int) (model.Task, error) {
	task, err := s.repo.Task.GetByID(ctx, taskId)
	if err != nil {
		s.logger.Error(err)
		if errors.Is(err, db.ErrNotFound) {
			return model.Task{}, service_errors.ErrNotFound
		}
		return model.Task{}, service_errors.ErrServiceNotAvailable
	}

	if task.GoalID != goalId {
		s.logger.Error("Error (GetByID): permission denied -> Task is not depended on Goal")
		return model.Task{}, service_errors.ErrNotFound
	}

	goal, err := s.repo.Goal.GetByID(ctx, task.GoalID)
	if err != nil {
		s.logger.Error(err)
		var finalErr = service_errors.ErrServiceNotAvailable
		if errors.Is(err, db.ErrNotFound) {
			finalErr = service_errors.ErrNotFound
		}
		return model.Task{}, finalErr
	}

	if goal.UserID != userId {
		s.logger.Error("Error (GetByID): permission denied -> Goal user id different")
		return model.Task{}, service_errors.ErrPermissionDenied
	}

	return task, nil
}

func (s Service) UpdateByID(ctx context.Context, userId, goalId, taskId int, task dto.UpdateTask) error {
	if err := task.Validate(); err != nil {
		return service_errors.ErrInvalidCredentials
	}

	_, err := s.GetByID(ctx, userId, goalId, taskId)
	if err != nil {
		s.logger.Errorf("Error (UpdateByID): %v", err)
		return err
	}

	err = s.repo.Task.UpdateByID(ctx, taskId, task)
	if err != nil {
		s.logger.Error(err)
		if errors.Is(err, db.ErrConflict) {
			return service_errors.ErrAlreadyExists
		}
		return service_errors.ErrServiceNotAvailable
	}
	return nil
}

func (s Service) DeleteByID(ctx context.Context, userId, goalId, taskId int) error {
	_, err := s.GetByID(ctx, userId, goalId, taskId)
	if err != nil {
		s.logger.Errorf("Error (DeleteByID): %v", err)
		return err
	}

	err = s.repo.Task.DeleteByID(ctx, taskId)
	if err != nil {
		s.logger.Error(err)
		return service_errors.ErrServiceNotAvailable
	}
	return nil
}
