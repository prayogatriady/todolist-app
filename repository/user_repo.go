package repository

import (
	"github.com/prayogatriady/todolist-app/model/entity"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(userEntity entity.User) (entity.User, error)
	GetUser(userID int64) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	GetUserByUsernamePassword(username string, password string) (entity.User, error)
	UpdateUser(userID int64, updateUser entity.User) (entity.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(userEntity entity.User) (entity.User, error) {
	if err := r.DB.Create(&userEntity).Error; err != nil {
		return userEntity, err
	}
	return userEntity, nil
}

func (r *UserRepo) GetUser(userID int64) (entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id =?", userID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	if err := r.DB.Where("username =?", username).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsernamePassword(username string, password string) (entity.User, error) {
	var user entity.User
	if err := r.DB.Where("username =? AND password =?", username, password).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) UpdateUser(userID int64, updateUser entity.User) (entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id =?", userID).Updates(&updateUser).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
