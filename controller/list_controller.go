package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/todolist-app/middleware"
	"github.com/prayogatriady/todolist-app/model/web"
	"github.com/prayogatriady/todolist-app/service"
)

type ListContInterface interface {
	CreateTodolist(c *gin.Context)
	GetList(c *gin.Context)
	EditList(c *gin.Context)
	DeleteList(c *gin.Context)
}

type ListCont struct {
	ListService service.ListServInterface
}

func NewListCont(listService service.ListServInterface) *ListCont {
	return &ListCont{
		ListService: listService,
	}
}

func (uc *ListCont) CreateTodolist(c *gin.Context) {
	// get payload from token
	userID, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// Get request body
	var listRequest web.ListRequest
	if err := c.BindJSON(&listRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	listResponse, err := uc.ListService.CreateTodolist(userID, listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Todolist created",
		"body":    listResponse,
	})
}

func (uc *ListCont) GetList(c *gin.Context) {
	// get payload from token
	userID, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	listsResponse, err := uc.ListService.GetLists(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Todolist retrieved",
		"body":    listsResponse,
	})
}

func (uc *ListCont) EditList(c *gin.Context) {
	// get payload from token
	userID, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// Get request body
	var listRequest web.ListRequest
	if err := c.BindJSON(&listRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	// Get todolistID as a Param
	stringTodolistID := c.Param("todolistID")
	todolistID, err := strconv.ParseInt(stringTodolistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	listResponse, err := uc.ListService.EditList(todolistID, userID, listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Todolist updated",
		"body":    listResponse,
	})
}

func (uc *ListCont) DeleteList(c *gin.Context) {
	// get payload from token
	userID, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}
	// Get todolistID as a Param
	stringTodolistID := c.Param("todolistID")
	todolistID, err := strconv.ParseInt(stringTodolistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	if err := uc.ListService.DeleteList(todolistID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Todolist deleted",
	})
}
