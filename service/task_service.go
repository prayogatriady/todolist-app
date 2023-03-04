package service

import (
	"github.com/prayogatriady/todolist-app/model/entity"
	"github.com/prayogatriady/todolist-app/model/web"
	"github.com/prayogatriady/todolist-app/repository"
)

type TaskServInterface interface {
	CreateTask(listID int64, userID int64, taskRequest web.TaskRequest) (web.TaskResponse, error)
	GetTasks(listID int64, userID int64) ([]web.TaskResponse, error)
	EditTask(taskID int64, listID int64, userID int64, taskRequest web.TaskRequest) (web.TaskResponse, error)
	DeleteTask(taskID int64, listID int64, userID int64) error
}

type TaskServ struct {
	TaskRepo repository.TaskRepoInterface
}

func NewTaskServ(taskRepo repository.TaskRepoInterface) TaskServInterface {
	return &TaskServ{
		TaskRepo: taskRepo,
	}
}

func (ts *TaskServ) CreateTask(listID int64, userID int64, taskRequest web.TaskRequest) (web.TaskResponse, error) {
	var taskEntity entity.Task
	taskEntity = entity.Task{
		Description: taskRequest.Description,
		Completed:   taskRequest.Completed,
		ListID:      listID,
		UserID:      userID,
	}

	taskEntity, err := ts.TaskRepo.CreateTask(taskEntity)
	if err != nil {
		return web.TaskResponse{}, err
	}

	var taskResponse web.TaskResponse
	taskResponse = web.TaskResponse{
		ID:          taskEntity.ID,
		Description: taskEntity.Description,
		Completed:   taskEntity.Completed,
		ListID:      taskEntity.ListID,
		UserID:      taskEntity.UserID,
	}

	return taskResponse, nil
}

func (ts *TaskServ) GetTasks(listID int64, userID int64) ([]web.TaskResponse, error) {
	taskEntity, err := ts.TaskRepo.GetTasks(listID, userID)
	if err != nil {
		return []web.TaskResponse{}, err
	}

	var tasksResponse []web.TaskResponse

	for _, task := range taskEntity {
		l := web.TaskResponse{
			ID:          task.ID,
			Description: task.Description,
			Completed:   task.Completed,
			ListID:      task.ListID,
			UserID:      task.UserID,
		}
		tasksResponse = append(tasksResponse, l)
	}

	return tasksResponse, nil
}

func (ts *TaskServ) EditTask(taskID int64, listID int64, userID int64, taskRequest web.TaskRequest) (web.TaskResponse, error) {
	var taskEntity entity.Task
	taskEntity = entity.Task{
		Description: taskRequest.Description,
		Completed:   taskRequest.Completed,
		ListID:      listID,
		UserID:      userID,
	}
	taskEntity, err := ts.TaskRepo.UpdateTask(taskID, listID, userID, taskEntity)
	if err != nil {
		return web.TaskResponse{}, err
	}

	var taskResponse web.TaskResponse
	taskResponse = web.TaskResponse{
		ID:          taskEntity.ID,
		Description: taskEntity.Description,
		Completed:   taskEntity.Completed,
		ListID:      taskEntity.ListID,
		UserID:      taskEntity.UserID,
	}

	return taskResponse, nil
}

func (ts *TaskServ) DeleteTask(taskID int64, listID int64, userID int64) error {
	if err := ts.TaskRepo.DeleteTask(taskID, listID, userID); err != nil {
		return err
	}
	return nil
}
