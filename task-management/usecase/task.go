package usecase

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"sort"
	"strings"
	"task-management-api/task-management/entities"
	"task-management-api/task-management/models"
	repo "task-management-api/task-management/repositories"
)

type task struct {
	taskRepo repo.TaskRepository
	logger   *zap.SugaredLogger
}

func NewTaskUsecase(taskRepo repo.TaskRepository, logger *zap.SugaredLogger) TaskUsecase {

	return &task{
		taskRepo: taskRepo,
		logger:   logger,
	}
}

func (u *task) CreateTask(req *models.TaskRequest) error {
	logger := u.logger.With(zap.String("Usecase", "CreateTask"))
	dataTask := entities.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      entities.StatusTodo,
	}

	if strings.TrimSpace(req.Title) == "" || strings.TrimSpace(req.Description) == "" {
		errs := errors.New("request body are require title and description")
		u.logger.Error("Request body are require title and description.", zap.Error(errs))
		fmt.Println(2, errs)
		return errs
	}

	logger.Info("Create with payload: ", zap.Any("payload", dataTask))
	err := u.taskRepo.CreateTask(dataTask)
	if err != nil {
		u.logger.Error("Can not create task", zap.Error(err))
		return err
	}

	return nil
}

func (u *task) GetTaskByID(id int) (*models.TaskResponse, error) {
	logger := u.logger.With(zap.String("Usecase", "GetTaskByID"))

	logger.Info("Request Task ID", zap.Int("id", id))

	taskData, err := u.taskRepo.GetTaskByID(id)
	if err != nil {
		u.logger.Error("Can not get task", zap.Error(err))
		return nil, err
	}

	respTask := new(models.TaskResponse)
	if err := mapstructure.Decode(taskData, &respTask); err != nil {
		u.logger.Error("Can map structure ", zap.Error(err))
		return nil, err
	}

	return respTask, nil
}

func (u *task) UpdateTask(id int, req *models.TaskRequest) error {
	logger := u.logger.With(zap.String("Usecase", "UpdateTask"))

	logger.Info("Request Task ID", zap.Int("id", id), zap.Any("reqBody", req))

	task, err := u.taskRepo.GetTaskByID(id)
	if err != nil {
		u.logger.Error("Task is not exist", zap.Error(err))
		return err
	}

	if strings.TrimSpace(req.Title) != "" {
		task.Title = req.Title
	}

	if strings.TrimSpace(req.Description) != "" {
		task.Description = req.Description
	}

	if err := u.taskRepo.UpdateTask(id, *task); err != nil {
		u.logger.Error("Failed to update task", zap.Error(err))
		return err
	}

	return nil
}

func (u *task) UpdateTaskStatus(id int, req *models.UpdateStatusRequest) error {
	logger := u.logger.With(zap.String("Usecase", "UpdateTaskStatus"))
	logger.Info("Request Task ID", zap.Int("id", id), zap.Any("reqBody", req))

	if err := u.taskRepo.UpdateTaskStatus(id, req.Status); err != nil {
		u.logger.Error("Failed to update task status", zap.Error(err))
		return err
	}
	return nil
}

func (u *task) DeleteTaskByID(id int) error {
	logger := u.logger.With(zap.String("Usecase", "DeleteTaskByID"))
	logger.Info("Request Task ID", zap.Int("id", id))

	if err := u.taskRepo.DeleteTask(id); err != nil {
		u.logger.Error("Failed to delete task", zap.Error(err))
		return err
	}
	return nil
}

func (u *task) ListTasks() (*[]models.TaskResponse, error) {
	logger := u.logger.With(zap.String("Usecase", "ListTasks"))
	logger.Info("Request List Tasks")

	tasks, err := u.taskRepo.ListTasks()
	if err != nil {
		u.logger.Error("Failed to delete task", zap.Error(err))
		return nil, err
	}

	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	taskData := new([]models.TaskResponse)
	if err := mapstructure.Decode(tasks, &taskData); err != nil {
		u.logger.Error("Can map structure ", zap.Error(err))
		return nil, err
	}

	return taskData, nil
}
