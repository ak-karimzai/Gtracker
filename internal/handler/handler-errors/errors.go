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
// @Description request error message
type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

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
