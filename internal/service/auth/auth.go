// Package auth provides authentication services.
package auth

import (
	"context"
	"errors"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository"
	service_errors "git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/service-errors"
	auth_token "git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/auth-token"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/db"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/util"
)

// Service provides authentication services.
type Service struct {
	repo       *repository.Repository
	tokenMaker auth_token.Maker
	logger     logger.Logger
}

// NewService creates a new instance of the authentication service.
func NewService(repo *repository.Repository, tokenMaker auth_token.Maker, logger logger.Logger) *Service {
	return &Service{repo: repo, tokenMaker: tokenMaker, logger: logger}
}

// SignUp registers a new user.
func (s Service) SignUp(ctx context.Context, up dto.SignUp) error {
	if err := up.Validate(); err != nil {
		s.logger.Error(err)
		return err
	}

	hashedPwd, err := util.HashPassword(up.Password)
	if err != nil {
		s.logger.Error(err)
		return service_errors.ErrServiceNotAvailable
	}
	up.Password = hashedPwd
	if err := s.repo.SignUp(ctx, up); err != nil {
		s.logger.Error(err)
		if errors.Is(err, db.ErrConflict) {
			return service_errors.ErrAlreadyExists
		}
		return service_errors.ErrServiceNotAvailable
	}

	return nil
}

// Login authenticates a user.
func (s Service) Login(ctx context.Context, in dto.Login) (*dto.LoginResponse, error) {
	if err := in.Validate(); err != nil {
		s.logger.Error(err)
		return nil, err
	}

	user, err := s.repo.GetUserByName(ctx, in.Username)
	if err != nil {
		s.logger.Error(err)
		if errors.Is(err, db.ErrNotFound) {
			return nil, service_errors.ErrNotFound
		}
		return nil, err
	}

	if err := util.CheckPwd(
		in.Password, user.Password); err != nil {
		s.logger.Error(err)
		return nil, service_errors.ErrServiceNotAvailable
	}

	token, err := s.tokenMaker.CreateToken(user.ID, user.Username)
	if err != nil {
		s.logger.Error(err)
		return nil, service_errors.ErrInvalidCredentials
	}

	return dto.NewLoginResponse(token, user), nil
}
