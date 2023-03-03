package service

import (
	"github.com/prayogatriady/todolist-app/middleware"
	"github.com/prayogatriady/todolist-app/model/entity"
	"github.com/prayogatriady/todolist-app/model/web"
	"github.com/prayogatriady/todolist-app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServInterface interface {
	CreateUser(userRequest web.UserSignupRequest) (web.UserResponse, error)
	Signin(userRequest web.UserSigninRequest) (web.UserTokenResponse, error)
}

type UserServ struct {
	UserRepo repository.UserRepoInterface
}

func NewUserServ(userRepo repository.UserRepoInterface) UserServInterface {
	return &UserServ{
		UserRepo: userRepo,
	}
}

func (us *UserServ) CreateUser(userRequest web.UserSignupRequest) (web.UserResponse, error) {
	// Generate hash password
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.UserResponse{}, err
	}

	var userEntity entity.User
	userEntity = entity.User{
		Username: userRequest.Username,
		Password: string(bytePassword),
	}

	userEntity, err = us.UserRepo.CreateUser(userEntity)
	if err != nil {
		return web.UserResponse{}, err
	}

	var userResponse web.UserResponse
	userResponse = web.UserResponse{
		ID:        userEntity.ID,
		Username:  userEntity.Username,
		Password:  userEntity.Password,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}

	return userResponse, nil
}

func (us *UserServ) Signin(userRequest web.UserSigninRequest) (web.UserTokenResponse, error) {
	userFound, err := us.UserRepo.GetUserByUsername(userRequest.Username)
	if err != nil {
		return web.UserTokenResponse{}, err
	}

	// compare found password from database and user input password
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userRequest.Password)); err != nil {
		return web.UserTokenResponse{}, err
	}

	user, err := us.UserRepo.GetUserByUsernamePassword(userFound.Username, userFound.Password)
	if err != nil {
		return web.UserTokenResponse{}, err
	}

	// create token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		return web.UserTokenResponse{}, err
	}

	return web.UserTokenResponse{Token: token}, nil
}
