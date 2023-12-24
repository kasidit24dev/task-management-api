package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"task-management-api/task-management/entities"
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

// @Version 1.0
// @Host localhost:3000
// @BasePath /api/v1
// @Title Example Response for Task Management API
// @Summary CreateTask
// @Description Create task
// @ID post-task
// @Accept json
// @Produce  json
// @Param request body models.TaskRequest true "Request Body for create task"
// @Success 201 {object} commonResponse "Success" example={"code":201,"message":"Success"}
// @Router /task [POST]
// @Tag Task-Management
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

// @Summary GetTaskByID
// @Description Get a task by ID
// @ID get-task-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} commonResponse{data=models.TaskResponse} "Success" example={"code":200,"message":"Success","data":{"id":2,"title":"2","description":"xxx","status":"To Do"}}
// @Router /task/{id} [GET]
// @Tag Tasks
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

// @Summary UpdateTask
// @Description Update a task by ID
// @ID update-task-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param request body models.TaskRequest true "Request Body for updating task"
// @Success 200 {object} commonResponse "Success" example={"code":200,"message":"Success"}
// @Router /task/{id} [PUT]
// @Tag Task Management
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

// @Summary UpdateTaskStatus
// @Description Update the status of a task by ID
// @ID update-task-status-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param status body models.UpdateStatusRequest true "New task status [To do, In Progress, Done]"
// @Success 200 {object} commonResponse "Success" example={"code":200,"message":"Success"}
// @Router /task/{id}/status [PATCH]
// @Tag Task Management
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
		logger.Error("Invalid request body", zap.Any("reqBody", fmt.Sprintf("%v", reqBody)), zap.Error(err))
		return responseError(c, 400, "Invalid request body", nil, err.Error())
	}

	if strings.TrimSpace(reqBody.Status) == "" {
		logger.Error("Invalid request body", zap.Any("reqBody", reqBody))
		return responseError(c, 400, "Invalid request payload", nil, "")
	}

	if reqBody.Status != entities.StatusTodo && reqBody.Status != entities.StatusInProgress && reqBody.Status != entities.StatusDone {
		logger.Error("Status not support")
		return responseError(c, 500, "Status not support", nil, "Please send a correct status")
	}

	if err := h.uc.UpdateTaskStatus(taskID, reqBody); err != nil {
		logger.Error("Internal Server Error", zap.Error(err))
		return responseError(c, 500, "Internal Server Error", nil, err.Error())
	}

	logger.Info("Update task status successfully ..")
	return responseOK(c, 200, "Success", nil)
}

// @Summary DeleteTaskByID
// @Description Delete a task by ID
// @ID delete-task-by-id
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} commonResponse "Success" example={"code":200,"message":"Success"}
// @Router /task/{id} [DELETE]
// @Tag Task Management
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

// @Summary ListTasks
// @Description Retrieve a list of tasks
// @ID list-tasks
// @Produce  json
// @Success 200 {object} commonResponse{data=[]models.TaskResponse} "Success" example={"code":200,"message":"Success","data":[{"id":1,"title":"2","description":"xxx","status":"To Do"},{"id":2,"title":"2","description":"xxx","status":"To Do"}]}
// @Router /tasks [GET]
// @Tag Task Management
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
