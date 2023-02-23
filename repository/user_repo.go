package repository

import (
	model "github.com/prayogatriady/todolist-app/model/entities"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(userID int64) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserByUsernamePassword(username string, password string) (model.User, error)
	UpdateUser(userID int64, updateUser model.User) (model.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(user model.User) (model.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) GetUser(userID int64) (model.User, error) {
	var user model.User
	if err := r.DB.Where("id =?", userID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("username =?", username).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsernamePassword(username string, password string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("username =? AND password =?", username, password).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) UpdateUser(userID int64, updateUser model.User) (model.User, error) {
	var user model.User
	if err := r.DB.Where("id =?", userID).Updates(&updateUser).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
