// Package goal provides services related to user goals.
package goal

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

// Service provides methods for managing user goals.
type Service struct {
	repo   *repository.Repository
	logger logger.Logger
}

// NewService creates a new instance of the goal service.
func NewService(repo *repository.Repository, logger logger.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

// Create creates a new goal for a user.
func (s Service) Create(ctx context.Context, userId int, goal dto.CreateGoal) (model.Goal, error) {
	if err := goal.Validate(); err != nil {
		s.logger.Error(err)
		return model.Goal{}, service_errors.ErrInvalidCredentials
	}

	id, err := s.repo.Goal.Create(ctx, userId, goal)
	if err != nil {
		s.logger.Error(err)
		var finalErr = service_errors.ErrServiceNotAvailable
		if errors.Is(err, db.ErrConflict) {
			finalErr = service_errors.ErrAlreadyExists
		}
		return model.Goal{}, finalErr
	}

	goalFromDB, err := s.repo.Goal.GetByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		var finalErr = service_errors.ErrServiceNotAvailable
		if errors.Is(err, db.ErrNotFound) {
			finalErr = service_errors.ErrNotFound
		}
		return model.Goal{}, finalErr
	}
	return goalFromDB, nil
}

// Get retrieves goals for a user based on provided parameters.
func (s Service) Get(ctx context.Context, userId int, listParams dto.ListParams) ([]model.Goal, error) {
	if err := listParams.Validate(); err != nil {
		s.logger.Error(err)
		return []model.Goal{}, service_errors.ErrInvalidCredentials
	}

	goals, err := s.repo.Goal.Get(ctx, userId, listParams)
	if err != nil {
		s.logger.Error(err)
		var finalErr = service_errors.ErrServiceNotAvailable
		if errors.Is(err, db.ErrNotFound) {
			finalErr = service_errors.ErrNotFound
		}
		return []model.Goal{}, finalErr
	}
	return goals, nil
}

// GetByID retrieves a goal by its ID and ensures the user has permission to access it.
func (s Service) GetByID(ctx context.Context, userId, goalId int) (model.Goal, error) {
	goal, err := s.repo.Goal.GetByID(ctx, goalId)
	if err != nil {
		s.logger.Error(err)
		if errors.Is(err, db.ErrNotFound) {
			return model.Goal{}, service_errors.ErrNotFound
		}
		return model.Goal{}, service_errors.ErrServiceNotAvailable
	}

	if goal.UserID != userId {
		return model.Goal{}, service_errors.ErrPermissionDenied
	}

	return goal, nil
}

// UpdateByID updates a goal by its ID.
func (s Service) UpdateByID(ctx context.Context, userId, goalId int, update dto.UpdateGoal) error {
	if err := update.Validate(); err != nil {
		s.logger.Error(err)
		return service_errors.ErrInvalidCredentials
	}

	_, err := s.GetByID(ctx, userId, goalId)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	err = s.repo.Goal.UpdateByID(ctx, goalId, update)
	if err != nil {
		s.logger.Error(err)
		if errors.Is(err, db.ErrConflict) {
			return service_errors.ErrAlreadyExists
		}
		return service_errors.ErrServiceNotAvailable
	}
	return nil
}

// DeleteByID deletes a goal by its ID.
func (s Service) DeleteByID(ctx context.Context, userId, goalId int) error {
	_, err := s.GetByID(ctx, userId, goalId)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	err = s.repo.Goal.DeleteByID(ctx, goalId)
	if err != nil {
		s.logger.Error(err)
		return service_errors.ErrServiceNotAvailable
	}
	return nil
}
