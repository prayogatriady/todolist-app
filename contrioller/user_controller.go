package controller

import (
	"github.com/gin-gonic/gin"
	model "github.com/prayogatriady/todolist-app/model/json"
	"github.com/prayogatriady/todolist-app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserContInterface interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)

	Profile(c *gin.Context)
	EditProfile(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserCont struct {
	// UserRepo repository.UserRepoInterface
	UserService service.UserContInterface
}

func NewUserCont(userService service.UserContInterface) *UserCont {
	return &UserCont{
		UserService: userService,
	}
}

func (uc *UserCont) Signup(c *gin.Context) {
	// Get request body
	var userRequest model.UserSignupRequest
	if err := c.BindJSON(%userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}


	user, err := uc.UserService.CreateUser(userRequest)
}
