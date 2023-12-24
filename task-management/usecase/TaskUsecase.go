package usecase

import "task-management-api/task-management/models"

type TaskUsecase interface {
	CreateTask(in *models.TaskRequest) error
	GetTaskByID(id int) (*models.TaskResponse, error)
	UpdateTask(id int, req *models.TaskRequest) error
	UpdateTaskStatus(id int, req *models.UpdateStatusRequest) error
	DeleteTaskByID(id int) error
	ListTasks() (*[]models.TaskResponse, error)
}
