package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"task-management-api/config"
	"task-management-api/task-management/handlers"
	"task-management-api/task-management/repositories"
	"task-management-api/task-management/usecase"
)

type echoServer struct {
	app    *echo.Echo
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewEchoServer(cfg *config.Config, logger *zap.SugaredLogger) Server {
	return &echoServer{
		app:    echo.New(),
		cfg:    cfg,
		logger: logger,
	}
}

func (server *echoServer) Start() {
	server.app.Use(middleware.Logger())
	server.initTasksHandler()
	serverUrl := fmt.Sprintf(":%d", server.cfg.App.Port)
	server.app.Logger.Fatal(server.app.Start(serverUrl))
}

func (server *echoServer) initTasksHandler() {

	taskStore := repositories.NewTaskStore()
	uc := usecase.NewTaskUsecase(taskStore, server.logger)

	tasksHandler := handlers.NewTaskHandler(uc, server.logger)

	tasksRouter := server.app.Group("/api/v1")
	tasksRouter.GET("/swagger/*", echoSwagger.WrapHandler)

	tasksRouter.POST("/task", tasksHandler.CreateTask)
	tasksRouter.GET("/task/:id", tasksHandler.GetTaskByID)
	tasksRouter.PUT("/task/:id", tasksHandler.UpdateTask)
	tasksRouter.PATCH("/task/:id/status", tasksHandler.UpdateTaskStatus)
	tasksRouter.DELETE("/task/:id", tasksHandler.DeleteTaskByID)
	tasksRouter.GET("/tasks", tasksHandler.ListTasks)

}
