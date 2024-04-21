// Package handler_errors provides utilities to handle and convert service layer
// errors into HTTP responses. This includes predefined error messages and
// a function to create error responses.
package handler_errors

import (
	"errors"
	"net/http"

	service_errors "git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/service-errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrAlreadyExist      = errors.New("Already exists")
	ErrBadRequest        = errors.New("Bad request")
	ErrUnauthorized      = errors.New("Unauthorized")
	ErrForbidden         = errors.New("Access fobidden")
	ErrBadCredentials    = errors.New("Bad credentials")
	ErrServerUnavailable = errors.New("something wrong in server")
	ErrNotFound          = errors.New("Not found")
	ErrPermissionDenied  = errors.New("Permission denied")
)

// ErrorResponse
// @Description defines the structure of error messages sent to clients.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse sends a JSON formatted error response with a specific status code and message.
//
// Parameters:
//   - c: The context of the current HTTP request.
//   - statusCode: The HTTP status code to be sent in the response.
//   - message: The error message that will be included in the response.
//
// This function uses Gin's context to abort the current request and send a JSON response
// with the specified status and an error message.
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

// ParseServiceErrors maps service layer errors to HTTP status codes and returns appropriate errors.
// This helps in abstracting service layer errors and translating them into understandable HTTP responses.
//
// Parameters:
//   - err: The error encountered in the service layer.
//
// Returns:
//   - int: The appropriate HTTP status code based on the error.
//   - error: The corresponding error that should be communicated back to the client.
//
// ParseServiceErrors uses standard Go error checking to determine the type of error and maps it to
// the appropriate HTTP status code. It handles various predefined errors from the service layer and
// maps them to more specific handler errors where necessary.
func ParseServiceErrors(err error) (int, error) {
	var status = http.StatusBadRequest
	var finalErr = err

	if errors.Is(err, service_errors.ErrInvalidCredentials) {
		status = http.StatusBadRequest
		finalErr = ErrBadRequest
	} else if errors.Is(err, service_errors.ErrNotFound) {
		status = http.StatusNotFound
		finalErr = ErrNotFound
	} else if errors.Is(err, service_errors.ErrPermissionDenied) {
		status = http.StatusForbidden
		finalErr = ErrForbidden
	} else if errors.Is(err, service_errors.ErrServiceNotAvailable) {
		status = http.StatusInternalServerError
		finalErr = ErrServerUnavailable
	} else if errors.Is(err, service_errors.ErrAlreadyExists) {
		status = http.StatusConflict
		finalErr = ErrAlreadyExist
	}
	return status, finalErr
}
