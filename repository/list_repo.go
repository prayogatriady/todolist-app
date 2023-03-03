package repository

import (
	"github.com/prayogatriady/todolist-app/model/entity"
	"gorm.io/gorm"
)

type ListRepoInterface interface {
	CreateList(list entity.List) (entity.List, error)
	GetList(listID int64) (entity.List, error)
	GetListByUserID(userID string) (entity.List, error)
	UpdateList(listID int64, updatedList entity.List) (entity.List, error)
	DeleteList(listID int64) error
}

type ListRepo struct {
	db *gorm.DB
}

func NewListRepo(db *gorm.DB) ListRepoInterface {
	return &ListRepo{db: db}
}

func (r *ListRepo) CreateList(list entity.List) (entity.List, error) {
	if err := r.db.Create(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepo) GetList(listID int64) (entity.List, error) {
	var list entity.List
	if err := r.db.Where("id =?", listID).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepo) GetListByUserID(userID string) (entity.List, error) {
	var list entity.List
	if err := r.db.Where("user_id =?", userID).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepo) UpdateList(listID int64, updatedList entity.List) (entity.List, error) {
	var list entity.List
	if err := r.db.Where("id =?", listID).Updates(&updatedList).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepo) DeleteList(listID int64) error {
	var list entity.List
	if err := r.db.Delete(&list, "id =?", listID).Error; err != nil {
		return err
	}
	return nil
}
