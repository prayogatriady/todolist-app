package repository

import (
	"github.com/prayogatriady/todolist-app/model/entity"
	"gorm.io/gorm"
)

type ListRepoInterface interface {
	CreateTodolist(list entity.List) (entity.List, error)
	GetListsByUserID(userID int64) ([]entity.List, error)
	UpdateList(listID int64, userID int64, updatedList entity.List) (entity.List, error)
	DeleteList(listID int64, userID int64) error
}

type ListRepo struct {
	db *gorm.DB
}

func NewListRepo(db *gorm.DB) ListRepoInterface {
	return &ListRepo{db: db}
}

func (r *ListRepo) CreateTodolist(list entity.List) (entity.List, error) {
	if err := r.db.Create(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepo) GetListsByUserID(userID int64) ([]entity.List, error) {
	var lists []entity.List
	if err := r.db.Where("user_id =?", userID).Find(&lists).Error; err != nil {
		return lists, err
	}
	return lists, nil
}

func (r *ListRepo) UpdateList(listID int64, userID int64, updatedList entity.List) (entity.List, error) {
	var list entity.List
	if err := r.db.Where("id =? AND user_id =?", listID, userID).Updates(&updatedList).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepo) DeleteList(listID int64, userID int64) error {
	var list entity.List
	if err := r.db.Delete(&list, "id =? AND user_id =?", listID, userID).Error; err != nil {
		return err
	}
	return nil
}
