// Package service_errors provides commonly used error variables for service-related errors.
package service_errors

import "errors"

// ErrInvalidCredentials is returned when the provided credentials are invalid.
var ErrInvalidCredentials = errors.New("invalid login credentials")

// ErrServiceNotAvailable is returned when a service is not available.
var ErrServiceNotAvailable = errors.New("service not available")

// ErrAlreadyExists is returned when an entity already exists.
var ErrAlreadyExists = errors.New("already exists")

// ErrNotFound is returned when an entity is not found.
var ErrNotFound = errors.New("not found")

// ErrPermissionDenied is returned when access to an entity is denied.
var ErrPermissionDenied = errors.New("access to entity denied")
