// Package auth provides authentication handling functionalities for the Gtracker application.
// It includes handling user sign-up and login processes, leveraging the underlying services
// to validate and persist user data.
package auth

import (
	"errors"
	"net/http"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	handler_errors "git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler/handler-errors"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service"
	service_errors "git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service/service-errors"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Handler structures necessary dependencies for handling authentication requests.
type Handler struct {
	service *service.Service
	logger  logger.Logger
}

// NewHandler creates a new authentication handler instance.
// It requires a service and a logger as dependencies to process the authentication requests.
//
// Parameters:
// - service: a reference to a Service that provides the business logic for authentication.
// - logger:  a Logger instance used for logging errors or other important system information.
//
// Returns:
// - *Handler: returns a pointer to a new Handler instance.
func NewHandler(service *service.Service, logger logger.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

// SignUp godoc
// @Summary      Signup
// @Description  Create an account in system
// @Tags         Auth
// @Accept 	  json
// @Produce 	  json
// @Param input body dto.SignUp true "sign up"
// @Success 201 {integer} integer 1
// @Failure 400 {object} handler_errors.ErrorResponse
// @Failure 409 {object} handler_errors.ErrorResponse
// @Router /auth/signup [post]
func (a *Handler) SignUp(ctx *gin.Context) {
	var request dto.SignUp
	if err := ctx.BindJSON(&request); err != nil {
		a.logger.Error(err)
		handler_errors.NewErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	if err := a.service.SignUp(ctx, request); err != nil {
		a.logger.Error(err)
		var status = http.StatusBadRequest
		var message = "invalid request"
		if errors.Is(err, service_errors.ErrAlreadyExists) {
			status = http.StatusConflict
			message = "user exists"
		}
		handler_errors.NewErrorResponse(ctx, status, message)
		return
	}
	ctx.Status(http.StatusCreated)
}

// Login godoc
// @Summary      Sign in
// @Description  Api for access to user privilege levels
// @Tags         Auth
// @Accept 	  	 json
// @Produce 	 json
// @Param input body dto.Login true "login"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} handler_errors.ErrorResponse
// @Failure 404 {object} handler_errors.ErrorResponse
// @Router /auth/login [post]
func (a *Handler) Login(ctx *gin.Context) {
	var request dto.Login
	if err := ctx.BindJSON(&request); err != nil {
		a.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			handler_errors.ErrBadRequest.Error())
		return
	}

	response, err := a.service.Login(ctx, request)
	if err != nil {
		a.logger.Error(err)
		status := http.StatusBadRequest
		message := handler_errors.ErrBadCredentials.Error()
		if errors.Is(err, service_errors.ErrNotFound) {
			status = http.StatusNotFound
			message = handler_errors.ErrNotFound.Error()
		} else if errors.Is(err, service_errors.ErrInvalidCredentials) {
			status = http.StatusBadRequest
			message = "incorrect credentials!"
		}
		handler_errors.NewErrorResponse(ctx, status, message)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
