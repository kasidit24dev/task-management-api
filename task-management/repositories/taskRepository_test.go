package repositories

import (
	"reflect"
	"sync"
	"task-management-api/task-management/entities"
	"testing"
)

func TestNewTaskStore(t *testing.T) {
	var tests []struct {
		name string
		want TaskRepository
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskStore_CreateTask(t *testing.T) {
	type fields struct {
		tasks  map[int]entities.Task
		lastID int
		mu     sync.RWMutex
	}
	type args struct {
		task entities.Task
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskStore{
				tasks:  tt.fields.tasks,
				lastID: tt.fields.lastID,
				mu:     tt.fields.mu,
			}
			if err := s.CreateTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_taskStore_DeleteTask(t *testing.T) {
	type fields struct {
		tasks  map[int]entities.Task
		lastID int
		mu     sync.RWMutex
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskStore{
				tasks:  tt.fields.tasks,
				lastID: tt.fields.lastID,
				mu:     tt.fields.mu,
			}
			if err := s.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_taskStore_GetTaskByID(t *testing.T) {
	type fields struct {
		tasks  map[int]entities.Task
		lastID int
		mu     sync.RWMutex
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Task
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskStore{
				tasks:  tt.fields.tasks,
				lastID: tt.fields.lastID,
				mu:     tt.fields.mu,
			}
			got, err := s.GetTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTaskByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskStore_ListTasks(t *testing.T) {
	type fields struct {
		tasks  map[int]entities.Task
		lastID int
		mu     sync.RWMutex
	}
	var tests []struct {
		name    string
		fields  fields
		want    []entities.Task
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskStore{
				tasks:  tt.fields.tasks,
				lastID: tt.fields.lastID,
				mu:     tt.fields.mu,
			}
			got, err := s.ListTasks()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskStore_UpdateTask(t *testing.T) {
	type fields struct {
		tasks  map[int]entities.Task
		lastID int
		mu     sync.RWMutex
	}
	type args struct {
		id         int
		updateTask entities.Task
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskStore{
				tasks:  tt.fields.tasks,
				lastID: tt.fields.lastID,
				mu:     tt.fields.mu,
			}
			if err := s.UpdateTask(tt.args.id, tt.args.updateTask); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_taskStore_UpdateTaskStatus(t *testing.T) {
	type fields struct {
		tasks  map[int]entities.Task
		lastID int
		mu     sync.RWMutex
	}
	type args struct {
		id     int
		status string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskStore{
				tasks:  tt.fields.tasks,
				lastID: tt.fields.lastID,
				mu:     tt.fields.mu,
			}
			if err := s.UpdateTaskStatus(tt.args.id, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTaskStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
