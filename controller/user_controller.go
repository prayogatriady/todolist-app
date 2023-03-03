package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/todolist-app/middleware"
	"github.com/prayogatriady/todolist-app/model/web"
	"github.com/prayogatriady/todolist-app/service"
)

type UserContInterface interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)

	Profile(c *gin.Context)
	EditProfile(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserCont struct {
	UserService service.UserServInterface
}

func NewUserCont(userService service.UserServInterface) *UserCont {
	return &UserCont{
		UserService: userService,
	}
}

func (uc *UserCont) Signup(c *gin.Context) {
	// Get request body
	var userRequest web.UserSignupRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	userResponse, err := uc.UserService.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User created",
		"body":    userResponse,
	})
}

func (uc *UserCont) Signin(c *gin.Context) {
	// Get request body
	var userRequest web.UserSigninRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	token, err := uc.UserService.Signin(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User signed in",
		"body":    token,
	})
}

func (uc *UserCont) Profile(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	userResponse, err := uc.UserService.Profile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User retrieved",
		"body":    userResponse,
	})
}

func (uc *UserCont) EditProfile(c *gin.Context) {
	// Get request body
	var userRequest web.UserEditRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	userResponse, err := uc.UserService.EditProfile(userId, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User updated",
		"body":    userResponse,
	})
}

func (uc *UserCont) DeleteUser(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	if err := uc.UserService.DeleteUser(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User deleted",
	})
}
