package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/prayogatriady/todolist-app/model/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(userEntity entity.User) (entity.User, error)
	GetUser(userID int64) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	GetUserByUsernamePassword(username string, password string) (entity.User, error)
	UpdateUser(userID int64, updateUser entity.User) (entity.User, error)
	DeleteUser(userID int64) error
}

type UserRepo struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewUserRepo(db *gorm.DB, rdb *redis.Client) UserRepoInterface {
	return &UserRepo{
		DB:  db,
		RDB: rdb,
	}
}

func (ur *UserRepo) CreateUser(userEntity entity.User) (entity.User, error) {
	if err := ur.DB.Create(&userEntity).Error; err != nil {
		return userEntity, err
	}
	return userEntity, nil
}

func (ur *UserRepo) GetUser(userID int64) (entity.User, error) {
	var user entity.User
	if err := ur.DB.Where("id =?", userID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepo) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User

	userRedis, err := ur.RDB.Get(context.Background(), username).Result()
	if err != nil {
		if err := ur.DB.Where("username =?", username).Find(&user).Error; err != nil {
			return user, err
		}

		userJson, err := json.Marshal(user)
		if err != nil {
			return user, err
		}
		ur.RDB.Set(context.Background(), username, userJson, time.Minute*10)

		return user, nil
	}

	err = json.Unmarshal([]byte(userRedis), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepo) GetUserByUsernamePassword(username string, password string) (entity.User, error) {
	var user entity.User
	if err := ur.DB.Where("username =? AND password =?", username, password).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepo) UpdateUser(userID int64, updateUser entity.User) (entity.User, error) {
	var user entity.User
	if err := ur.DB.Where("id =?", userID).Updates(&updateUser).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepo) DeleteUser(userID int64) error {
	if err := ur.DB.Where("id =?", userID).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil
}
