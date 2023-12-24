package repositories

import "task-management-api/task-management/entities"

type TaskRepository interface {
	CreateTask(task entities.Task) error
	UpdateTask(id int, updateTask entities.Task) error
	GetTaskByID(id int) (*entities.Task, error)
	UpdateTaskStatus(id int, status string) error
	DeleteTask(id int) error
	ListTasks() ([]entities.Task, error)
}
