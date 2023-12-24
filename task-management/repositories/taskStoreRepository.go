package repositories

import (
	"errors"
	"sync"
	"task-management-api/task-management/entities"
)

type taskStore struct {
	tasks  map[int]entities.Task
	lastID int
	mu     sync.RWMutex
}

func NewTaskStore() TaskRepository {

	return &taskStore{
		tasks:  make(map[int]entities.Task),
		lastID: 0,
	}
}

func (s *taskStore) CreateTask(task entities.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	task.ID = s.lastID
	s.tasks[task.ID] = task

	return nil
}

func (s *taskStore) GetTaskByID(id int) (*entities.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("[GetTask] task is not exists")
	}
	return &task, nil
}

func (s *taskStore) UpdateTask(id int, updateTask entities.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.tasks[id]; !exists {
		return errors.New("[UpdateTask] task is not exists")
	}

	updateTask.ID = id
	updateTask.Status = s.tasks[updateTask.ID].Status
	s.tasks[id] = updateTask

	return nil
}

func (s *taskStore) UpdateTaskStatus(id int, status string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, exists := s.tasks[id]; exists {
		task.Status = status
		s.tasks[id] = task

		return nil
	}
	return errors.New("[UpdateTaskStatus] Task is not exists")
}

func (s *taskStore) DeleteTask(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("[DeleteTask] Task is not exists")
	}

	delete(s.tasks, id)
	return nil
}

func (s *taskStore) ListTasks() ([]entities.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.tasks) == 0 {
		return nil, errors.New("[ListTasks] Tasks is empty")
	}
	tasks := make([]entities.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
