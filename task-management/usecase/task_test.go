package usecase

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"task-management-api/config"
	log "task-management-api/logger"
	"task-management-api/task-management/entities"
	"task-management-api/task-management/mocks"
	"task-management-api/task-management/models"
	"testing"
)

type taskUsecaseSuite struct {
	suite.Suite

	cfg    *config.Config
	logger *zap.SugaredLogger

	taskRepo *mocks.TaskRepository
	taskUC   TaskUsecase

	tasks map[int]entities.Task
}

func (suite *taskUsecaseSuite) SetupTest() {

	cfg := config.Config{App: config.App{
		Env:  "dev",
		Port: 3000,
	}}
	logger := log.InitLogger(&cfg)

	taskRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(taskRepo, logger)

	suite.taskRepo = taskRepo
	suite.taskUC = uc

	suite.tasks = make(map[int]entities.Task)
	suite.tasks[1] = entities.Task{
		ID:          1,
		Title:       "mock test",
		Description: "mock test",
		Status:      "To Do",
	}
}

func (suite *taskUsecaseSuite) TestCreateTaskSuccess() {
	reqData := models.TaskRequest{
		Title:       "task 01",
		Description: "task 01 description",
	}
	suite.taskRepo.On("CreateTask", entities.Task{
		Title:       "task 01",
		Description: "task 01 description",
		Status:      "To Do",
	}).Return(nil)

	err := suite.taskUC.CreateTask(&reqData)

	suite.Nil(err, "error")
	suite.taskRepo.AssertExpectations(suite.T())

}

func (suite *taskUsecaseSuite) TestCreateTaskPayloadFail() {
	reqData := models.TaskRequest{
		Description: "task 01 description",
	}

	err := suite.taskUC.CreateTask(&reqData)
	suite.Error(err, "request body are require title and description")
	suite.taskRepo.AssertExpectations(suite.T())

}

func (suite *taskUsecaseSuite) TestUpdateTaskSuccess() {
	id := 1
	reqData := &models.TaskRequest{
		Title:       "test",
		Description: "test",
	}

	mockTask := new(entities.Task)
	mockTask.Title = "test"
	mockTask.Description = "test"
	suite.taskRepo.On("GetTaskByID", id).Return(mockTask, nil)

	suite.taskRepo.On("UpdateTask", id, *mockTask).Return(nil)
	err := suite.taskUC.UpdateTask(id, reqData)
	suite.Nil(err, "error")
	suite.taskRepo.AssertExpectations(suite.T())
}
func (suite *taskUsecaseSuite) TestGetTaskByIDSuccess() {
	id := 1
	task := &entities.Task{
		ID:          1,
		Title:       "Test 01",
		Description: "Test 01",
		Status:      "To Do",
	}

	expect := &models.TaskResponse{
		ID:          1,
		Title:       "Test 01",
		Description: "Test 01",
		Status:      "To Do",
	}
	suite.taskRepo.On("GetTaskByID", id).Return(task, nil)

	result, err := suite.taskUC.GetTaskByID(id)
	suite.Nil(err, "[GetTask] task is not exists")
	suite.Equal(expect, result)
}

func (suite *taskUsecaseSuite) TestUpdateTaskStatusSuccess() {
	id := 1
	status := "Done"
	reqData := &models.UpdateStatusRequest{Status: status}

	suite.taskRepo.On("UpdateTaskStatus", id, status).Return(nil)

	err := suite.taskUC.UpdateTaskStatus(id, reqData)
	suite.Nil(err, "Failed to update task status")
}

func (suite *taskUsecaseSuite) TestDeleteTaskSuccess() {
	id := 1

	suite.taskRepo.On("DeleteTask", id).Return(nil)

	err := suite.taskUC.DeleteTaskByID(id)
	suite.Nil(err, "Failed to delete task")
}

func (suite *taskUsecaseSuite) TestGetListTasksSuccess() {

	dataTask := []entities.Task{{
		ID:          1,
		Title:       "test01",
		Description: "test01",
		Status:      "To Do",
	}, {
		ID:          2,
		Title:       "test02",
		Description: "test02",
		Status:      "To Do",
	}}

	expect := &[]models.TaskResponse{{
		ID:          1,
		Title:       "test01",
		Description: "test01",
		Status:      "To Do",
	}, {
		ID:          2,
		Title:       "test02",
		Description: "test02",
		Status:      "To Do",
	}}
	suite.taskRepo.On("ListTasks").Return(dataTask, nil)

	result, err := suite.taskUC.ListTasks()
	suite.Nil(err, "Failed to delete task")
	suite.Equal(expect, result)
}

func TestTask(t *testing.T) {
	suite.Run(t, new(taskUsecaseSuite))
}
