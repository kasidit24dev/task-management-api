package mocks

import (
	"github.com/stretchr/testify/mock"
	"task-management-api/task-management/entities"
)

type TaskRepository struct {
	mock.Mock
}

func (m *TaskRepository) CreateTask(task entities.Task) error {
	args := m.Called(task)

	var r0 error
	if rf, ok := args.Get(0).(func(entities.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (m *TaskRepository) UpdateTask(id int, updateTask entities.Task) error {
	args := m.Called(id, updateTask)

	var r0 error
	if rf, ok := args.Get(0).(func(int, entities.Task) error); ok {
		r0 = rf(id, updateTask)
	} else {
		r0 = args.Error(0)
	}
	return r0
}

func (m *TaskRepository) GetTaskByID(id int) (*entities.Task, error) {
	args := m.Called(id)

	var r0 *entities.Task
	if rf, ok := args.Get(0).(func() *entities.Task); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*entities.Task)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (m *TaskRepository) UpdateTaskStatus(id int, status string) error {

	args := m.Called(id, status)

	var r0 error
	if rf, ok := args.Get(0).(func(int, string) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = args.Error(0)
	}
	return r0
}

func (m *TaskRepository) DeleteTask(id int) error {
	args := m.Called(id)

	var r0 error
	if rf, ok := args.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = args.Error(0)
	}
	return r0
}

func (m *TaskRepository) ListTasks() ([]entities.Task, error) {
	args := m.Called()

	var r0 []entities.Task
	if rf, ok := args.Get(0).(func() []entities.Task); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]entities.Task)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
