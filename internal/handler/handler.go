// Package handler provides HTTP request handlers for different endpoints in the Gtracker application.
// It includes handlers for authentication, goal management, task management, and middleware initialization.
package handler

import (
	_ "git.iu7.bmstu.ru/ka19iu10/Gtracker/docs"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler/auth"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler/goal"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler/middleware"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler/task"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service"
	auth_token "git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/auth-token"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Handler struct defines the main handler for HTTP request routing in the Gtracker application.
// It includes handlers for goals, tasks, and authentication, as well as a token maker for user authentication.
type Handler struct {
	Goal  *goal.Handler    // Goal handles HTTP requests related to goal management.
	Task  *task.Handler    // Task handles HTTP requests related to task management.
	Auth  *auth.Handler    // Auth handles HTTP requests related to user authentication and authorization.
	Maker auth_token.Maker // Maker is used for token generation and verification.
}

// NewHandler creates a new instance of the Handler struct with the provided services, token maker, and logger.
//
// Parameters:
//   - services: A pointer to the service layer providing business logic for the application.
//   - tokenMaker: A token maker instance for user authentication.
//   - logger: A logger instance for logging events or errors.
//
// Returns:
//   - *Handler: A pointer to the newly created Handler instance.
func NewHandler(services *service.Service, tokenMaker auth_token.Maker, logger logger.Logger) *Handler {
	return &Handler{
		Goal:  goal.NewHandler(services, logger),
		Task:  task.NewHandler(services, logger),
		Auth:  auth.NewHandler(services, logger),
		Maker: tokenMaker,
	}
}

// InitRoutes initializes routes for HTTP request handling using Gin framework.
// It sets up routes for various endpoints including authentication, goal management, and task management.
//
// Parameters:
//   - basePath: The base path for the API endpoints.
//
// Returns:
//   - *gin.Engine: A Gin Engine instance with initialized routes.
func (handler *Handler) InitRoutes(basePath string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Cors())

	api := router.Group(basePath)
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth := api.Group("/auth")
		{
			auth.POST("/signup", handler.Auth.SignUp)
			auth.POST("/login", handler.Auth.Login)
		}

		goals := api.Group("/goals", middleware.UserAuthentication(handler.Maker))
		{
			goals.POST("/", handler.Goal.Create)
			goals.GET("/", handler.Goal.Get)
			goals.GET("/:id", handler.Goal.GetByID)
			goals.PATCH("/:id", handler.Goal.UpdateByID)
			goals.DELETE("/:id", handler.Goal.DeleteByID)
			goals.GET("/:id/tasks", handler.Task.Get)
			goals.POST("/:id/tasks", handler.Task.Create)
			goals.GET("/:id/tasks/:task_id", handler.Task.GetByID)
			goals.PUT("/:id/tasks/:task_id", handler.Task.UpdateByID)
			goals.DELETE("/:id/tasks/:task_id", handler.Task.DeleteByID)
		}
	}
	return router
}
