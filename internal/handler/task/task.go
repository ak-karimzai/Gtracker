package task

import (
	"fmt"
	"net/http"
	"strconv"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	handler_errors "git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler/handler-errors"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler/middleware"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	logger  logger.Logger
}

func NewHandler(service *service.Service, logger logger.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

// Create godoc
// @Security ApiKeyAuth
// @Summary      Create task for a goal by system user
// @Description  Create task request
// @Tags         Task
// @Accept 	  json
// @Produce 	  json
// @Param  id path int true "Goal id"
// @Param 	input body dto.CreateTask true "Create task request"
// @Success 201 {object} model.Task
// @Failure 400 {object} handler_errors.ErrorResponse
// @Failure 401 {object} handler_errors.ErrorResponse
// @Failure 403 {object} handler_errors.ErrorResponse
// @Failure 409 {object} handler_errors.ErrorResponse
// @Router /goals/{id}/tasks [post]
func (h Handler) Create(ctx *gin.Context) {
	var request dto.CreateTask

	payload, err := middleware.GetUserInfo(ctx)
	if err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusUnauthorized,
			handler_errors.ErrUnauthorized.Error())
		return
	}

	if err := ctx.ShouldBind(&request); err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			handler_errors.ErrBadRequest.Error())
		return
	}

	goalId, _ := strconv.Atoi(ctx.Param("id"))
	if goalId <= 0 {
		message := fmt.Sprintf("incorrect goal_id: %d", goalId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}

	task, err := h.service.Task.Create(
		ctx,
		payload.UserID,
		goalId,
		request)
	if err != nil {
		h.logger.Error(err)
		status, err := handler_errors.ParseServiceErrors(err)
		handler_errors.NewErrorResponse(
			ctx,
			status,
			err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

// Get godoc
// @Security ApiKeyAuth
// @Summary      Get user goal tasks
// @Description  Get a list of user goal tasks
// @Tags         Task
// @Accept 	  json
// @Produce 	  json
// @Param  	id 		path   int true "Goal id"
// @Param   page_id query  int   true  "number of the page"  minimum(1)  default(1)
// @Param   page_size query  int   true  "size of the page" minimum(5) minimum(20) default(5)
// @Success 200 {array} model.Task
// @Failure 400 {object} handler_errors.ErrorResponse
// @Failure 401 {object} handler_errors.ErrorResponse
// @Router /goals/{id}/tasks [get]
func (h Handler) Get(ctx *gin.Context) {
	var listParams dto.ListParams

	payload, err := middleware.GetUserInfo(ctx)
	if err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusUnauthorized,
			handler_errors.ErrUnauthorized.Error())
		return
	}

	if err = ctx.BindQuery(&listParams); err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			handler_errors.ErrBadRequest.Error())
		return
	}

	goalId, _ := strconv.Atoi(ctx.Param("id"))
	if goalId <= 0 {
		message := fmt.Sprintf("incorrect goal_id: %d", goalId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}

	tasks, err := h.service.Task.Get(ctx, payload.UserID, goalId, listParams)
	if err != nil {
		h.logger.Error(err)
		status, err := handler_errors.ParseServiceErrors(err)
		handler_errors.NewErrorResponse(
			ctx,
			status,
			err.Error())
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

// GetByID godoc
// @Security ApiKeyAuth
// @Summary      Get user goal task by id
// @Description  Get user goal task by id
// @Tags         Task
// @Accept 	  	  json
// @Produce 	  json
// @Param 	id path int true "Goal ID"
// @Param 	task_id path int true "Task ID"
// @Success 200 {object} model.Task
// @Failure 400 {object} handler_errors.ErrorResponse
// @Failure 401 {object} handler_errors.ErrorResponse
// @Failure 403 {object} handler_errors.ErrorResponse
// @Failure 404 {object} handler_errors.ErrorResponse
// @Router /goals/{id}/tasks/{task_id} [get]
func (h Handler) GetByID(ctx *gin.Context) {
	payload, err := middleware.GetUserInfo(ctx)
	if err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusUnauthorized,
			handler_errors.ErrUnauthorized.Error())
		return
	}

	goalId, _ := strconv.Atoi(ctx.Param("id"))
	if goalId <= 0 {
		message := fmt.Sprintf("incorrect goal_id: %d", goalId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}

	taskId, _ := strconv.Atoi(ctx.Param("task_id"))
	if taskId <= 0 {
		message := fmt.Sprintf("incorrect task_id: %d", taskId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}
	task, err := h.service.Task.GetByID(ctx, payload.UserID, goalId, taskId)
	if err != nil {
		h.logger.Error(err)
		status, err := handler_errors.ParseServiceErrors(err)
		handler_errors.NewErrorResponse(
			ctx,
			status,
			err.Error())
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// UpdateByID godoc
// @Security ApiKeyAuth
// @Summary      Update user goal task
// @Description  update user goal task by id
// @Tags         Task
// @Accept 	  json
// @Produce 	  json
// @Param 	id path int true "Goal ID"
// @Param 	task_id path int true "Task ID"
// @Param input body dto.UpdateTask true "Create goal"
// @Success 204 {integer}  1
// @Failure 400 {object} handler_errors.ErrorResponse
// @Failure 401 {object} handler_errors.ErrorResponse
// @Failure 403 {object} handler_errors.ErrorResponse
// @Failure 404 {object} handler_errors.ErrorResponse
// @Failure 409 {object} handler_errors.ErrorResponse
// @Router /goals/{id}/tasks/{task_id} [put]
func (h Handler) UpdateByID(ctx *gin.Context) {
	var request dto.UpdateTask

	payload, err := middleware.GetUserInfo(ctx)
	if err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusUnauthorized,
			handler_errors.ErrUnauthorized.Error())
		return
	}

	goalId, _ := strconv.Atoi(ctx.Param("id"))
	if goalId <= 0 {
		message := fmt.Sprintf("incorrect goal_id: %d", goalId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}

	taskId, _ := strconv.Atoi(ctx.Param("task_id"))
	if taskId <= 0 {
		message := fmt.Sprintf("incorrect task_id: %d", taskId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}
	if err := ctx.ShouldBind(&request); err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			handler_errors.ErrBadRequest.Error())
		return
	}

	err = h.service.Task.UpdateByID(ctx, payload.UserID, goalId, taskId, request)
	if err != nil {
		h.logger.Error(err)
		status, err := handler_errors.ParseServiceErrors(err)
		handler_errors.NewErrorResponse(
			ctx,
			status,
			err.Error())
		return
	}
	ctx.Status(http.StatusNoContent)
}

// DeleteByID godoc
// @Security ApiKeyAuth
// @Summary      Delete user goal task
// @Description  delete user goal by id
// @Tags         Task
// @Accept 	  json
// @Produce 	  json
// @Param 	id path int true "Goal ID"
// @Param 	task_id path int true "Task ID"
// @Success 200 {integer}  1
// @Failure 400 {object} handler_errors.ErrorResponse
// @Failure 401 {object} handler_errors.ErrorResponse
// @Failure 403 {object} handler_errors.ErrorResponse
// @Failure 404 {object} handler_errors.ErrorResponse
// @Router /goals/{id}/tasks/{task_id} [delete]
func (h Handler) DeleteByID(ctx *gin.Context) {
	payload, err := middleware.GetUserInfo(ctx)
	if err != nil {
		h.logger.Error(err)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusUnauthorized,
			handler_errors.ErrUnauthorized.Error())
		return
	}

	goalId, _ := strconv.Atoi(ctx.Param("id"))
	if goalId <= 0 {
		message := fmt.Sprintf("incorrect goal_id: %d", goalId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}

	taskId, _ := strconv.Atoi(ctx.Param("task_id"))
	if taskId <= 0 {
		message := fmt.Sprintf("incorrect task_id: %d", taskId)
		h.logger.Error(message)
		handler_errors.NewErrorResponse(
			ctx,
			http.StatusBadRequest,
			message)
		return
	}

	if err := h.service.Task.DeleteByID(ctx, payload.UserID, goalId, taskId); err != nil {
		h.logger.Error(err)
		status, err := handler_errors.ParseServiceErrors(err)
		handler_errors.NewErrorResponse(
			ctx,
			status,
			err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}
