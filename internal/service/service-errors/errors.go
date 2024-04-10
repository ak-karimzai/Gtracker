package service_errors

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid login credentials")
	ErrServiceNotAvailable = errors.New("service not available")
	ErrAlreadyExists       = errors.New("already exist")
	ErrNotFound            = errors.New("not found")
	ErrPermissionDenied    = errors.New("access to entity denied")
)
