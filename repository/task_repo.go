package repository

import (
	model "github.com/prayogatriady/todolist-app/model/entities"
	"gorm.io/gorm"
)

type TaskRepoInterface interface {
	CreateTask(task model.Task) (model.Task, error)
	GetTask(taskID int64) (model.Task, error)
	GetTaskByListID(listID string) (model.Task, error)
	UpdateTask(taskID int64, updatedTask model.Task) (model.Task, error)
	DeleteTask(taskID int64) error
}

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepoInterface {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(task model.Task) (model.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) GetTask(taskID int64) (model.Task, error) {
	var task model.Task
	if err := r.db.Where("id =?", taskID).Find(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) GetTaskByListID(listID string) (model.Task, error) {
	var task model.Task
	if err := r.db.Where("user_id =?", listID).Find(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) UpdateTask(taskID int64, updatedTask model.Task) (model.Task, error) {
	var task model.Task
	if err := r.db.Where("id =?", taskID).Updates(&updatedTask).Find(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) DeleteTask(taskID int64) error {
	var task model.Task
	if err := r.db.Delete(&task, "id =?", taskID).Error; err != nil {
		return err
	}
	return nil
}
