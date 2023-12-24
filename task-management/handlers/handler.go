package handlers

import "github.com/labstack/echo/v4"

type Handler interface {
	CreateTask(c echo.Context) error
	GetTaskByID(c echo.Context) error
	UpdateTask(c echo.Context) error
	UpdateTaskStatus(c echo.Context) error
	DeleteTaskByID(c echo.Context) error
	ListTasks(c echo.Context) error
}
