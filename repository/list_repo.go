package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/prayogatriady/todolist-app/model/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ListRepoInterface interface {
	CreateTodolist(list entity.List) (entity.List, error)
	GetListsByUserID(userID int64) ([]entity.List, error)
	UpdateList(listID int64, userID int64, updatedList entity.List) (entity.List, error)
	DeleteList(listID int64, userID int64) error
}

type ListRepo struct {
	db  *gorm.DB
	RDB *redis.Client
}

func NewListRepo(db *gorm.DB, rdb *redis.Client) ListRepoInterface {
	return &ListRepo{
		db:  db,
		RDB: rdb,
	}
}

func (r *ListRepo) CreateTodolist(list entity.List) (entity.List, error) {
	if err := r.db.Create(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepo) GetListsByUserID(userID int64) ([]entity.List, error) {
	var lists []entity.List

	listsKey := "list" + strconv.Itoa(int(userID))
	listsRedis, err := r.RDB.Get(context.Background(), listsKey).Result()
	if err != nil {
		if err := r.db.Where("user_id =?", userID).Find(&lists).Error; err != nil {
			return lists, err
		}

		listsJson, err := json.Marshal(lists)
		if err != nil {
			return lists, err
		}
		r.RDB.Set(context.Background(), listsKey, listsJson, time.Minute*10)

		return lists, nil
	}

	err = json.Unmarshal([]byte(listsRedis), &lists)
	if err != nil {
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
