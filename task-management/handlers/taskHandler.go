package handlers

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"task-management-api/task-management/models"
	"task-management-api/task-management/usecase"
)

type taskHandler struct {
	uc     usecase.TaskUsecase
	logger *zap.SugaredLogger
}

func NewTaskHandler(uc usecase.TaskUsecase, logger *zap.SugaredLogger) Handler {

	return taskHandler{
		uc:     uc,
		logger: logger,
	}
}

func (h taskHandler) CreateTask(c echo.Context) error {
	logger := h.logger.With(zap.String("Handler", "CreateTask"))
	logger.Info("Starting create task...")

	reqBody := new(models.TaskRequest)
	if err := c.Bind(reqBody); err != nil {
		logger.Error("Invalid request body", zap.Any("reqBody", reqBody), zap.Error(err))
		return responseError(c, 400, "Invalid request body", nil, err.Error())
	}

	if strings.TrimSpace(reqBody.Title) == "" || strings.TrimSpace(reqBody.Description) == "" {
		logger.Error("Invalid request body", zap.Any("reqBody", reqBody))
		return responseError(c, 400, "Invalid request payload", nil, "")
	}

	err := h.uc.CreateTask(reqBody)
	if err != nil {
		logger.Error("Can not create task", zap.Error(err))
		return responseError(c, 500, "Can not create task", nil, err.Error())
	}
	logger.Info("Create task successfully ..")
	return responseOK(c, 200, "Success", nil)
}

func (h taskHandler) GetTaskByID(c echo.Context) error {
	logger := h.logger.With(zap.String("Handler", "GetTaskByID"))
	logger.Info("Starting get task...")

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid task ID", zap.Error(err))
		return responseError(c, 400, "Invalid task ID", nil, err.Error())
	}

	task, err := h.uc.GetTaskByID(taskID)
	if err != nil {
		logger.Error("Can not create task", zap.Error(err))
		return responseError(c, 400, "Task not found", nil, err.Error())
	}

	logger.Info("Get task successfully ..")
	return responseOK(c, 200, "Success", task)
}

func (h taskHandler) UpdateTask(c echo.Context) error {

	logger := h.logger.With(zap.String("Handler", "UpdateTask"))
	logger.Info("Starting update task...")
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid task ID", zap.Error(err))
		return responseError(c, 400, "Invalid task ID", nil, err.Error())
	}

	reqBody := new(models.TaskRequest)
	if err := c.Bind(reqBody); err != nil {
		logger.Error("Invalid request body", zap.Any("reqBody", reqBody), zap.Error(err))
		return responseError(c, 400, "Invalid request body", nil, err.Error())
	}

	if strings.TrimSpace(reqBody.Title) == "" || strings.TrimSpace(reqBody.Description) == "" {
		logger.Error("Invalid request body", zap.Any("reqBody", reqBody))
		return responseError(c, 400, "Invalid request payload", nil, "")
	}

	if err := h.uc.UpdateTask(taskID, reqBody); err != nil {
		logger.Error("Internal Server Error", zap.Error(err))
		return responseError(c, 500, "Internal Server Error", nil, err.Error())
	}
	logger.Info("Update task successfully ..")
	return responseOK(c, 200, "Success", nil)
}

func (h taskHandler) UpdateTaskStatus(c echo.Context) error {
	logger := h.logger.With(zap.String("Handler", "UpdateTaskStatus"))
	logger.Info("Starting update task status ...")

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid task ID", zap.Error(err))
		return responseError(c, 400, "Invalid task ID", nil, err.Error())
	}

	reqBody := new(models.UpdateStatusRequest)
	if err := c.Bind(reqBody); err != nil {
		logger.Error("Invalid request body", zap.Any("reqBody", reqBody), zap.Error(err))
		return responseError(c, 400, "Invalid request body", nil, err.Error())
	}

	if strings.TrimSpace(reqBody.Status) == "" {
		logger.Error("Invalid request body", zap.Any("reqBody", reqBody))
		return responseError(c, 400, "Invalid request payload", nil, "")
	}

	if err := h.uc.UpdateTaskStatus(taskID, reqBody); err != nil {
		logger.Error("Internal Server Error", zap.Error(err))
		return responseError(c, 500, "Internal Server Error", nil, err.Error())
	}

	logger.Info("Update task status successfully ..")
	return responseOK(c, 200, "Success", nil)
}

func (h taskHandler) DeleteTaskByID(c echo.Context) error {
	logger := h.logger.With(zap.String("Handler", "DeleteTaskByID"))
	logger.Info("Starting delete task by id...")

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid task ID", zap.Error(err))
		return responseError(c, 400, "Invalid task ID", nil, err.Error())
	}

	if err := h.uc.DeleteTaskByID(taskID); err != nil {
		logger.Error("Internal Server Error", zap.Error(err))
		return responseError(c, 500, "Internal Server Error", nil, err.Error())
	}

	logger.Info("Delete task successfully ..")
	return responseOK(c, 200, "Success", nil)
}

func (h taskHandler) ListTasks(c echo.Context) error {
	logger := h.logger.With(zap.String("Handler", "ListTasks"))
	logger.Info("Starting get list tasks...")

	tasks, err := h.uc.ListTasks()
	if err != nil {
		logger.Error("Internal Server Error", zap.Error(err))
		return responseError(c, 500, "Internal Server Error", nil, err.Error())
	}

	logger.Info("Delete task successfully ..")
	return responseOK(c, 200, "Success", tasks)
}
