package repository

import (
	"github.com/prayogatriady/todolist-app/model/entity"
	"gorm.io/gorm"
)

type TaskRepoInterface interface {
	CreateTask(task entity.Task) (entity.Task, error)
	GetTasks(listID int64, userID int64) ([]entity.Task, error)
	UpdateTask(taskID int64, listID int64, userID int64, updatedTask entity.Task) (entity.Task, error)
	DeleteTask(taskID int64, listID int64, userID int64) error
}

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepoInterface {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(task entity.Task) (entity.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) GetTasks(listID int64, userID int64) ([]entity.Task, error) {
	var tasks []entity.Task
	if err := r.db.Where("list_id =? AND user_id =?", listID, userID).Find(&tasks).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *TaskRepo) UpdateTask(taskID int64, listID int64, userID int64, updatedTask entity.Task) (entity.Task, error) {
	var task entity.Task
	if err := r.db.Where("id =? AND list_id =? AND user_id =?", taskID, listID, userID).Updates(&updatedTask).Find(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) DeleteTask(taskID int64, listID int64, userID int64) error {
	var task entity.Task
	if err := r.db.Delete(&task, "id =? AND list_id =? AND user_id =?", taskID, listID, userID).Error; err != nil {
		return err
	}
	return nil
}
