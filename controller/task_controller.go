package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/todolist-app/middleware"
	"github.com/prayogatriady/todolist-app/model/web"
	"github.com/prayogatriady/todolist-app/service"
)

type TaskContInterface interface {
	CreateTask(c *gin.Context)
	GetTasks(c *gin.Context)
	EditTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type TaskCont struct {
	TaskService service.TaskServInterface
}

func NewTaskCont(taskService service.TaskServInterface) *TaskCont {
	return &TaskCont{
		TaskService: taskService,
	}
}

func (uc *TaskCont) CreateTask(c *gin.Context) {
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
	var taskRequest web.TaskRequest
	if err := c.BindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	// Get query param
	stringTodolistID := c.Query("todolistID")
	todolistID, err := strconv.ParseInt(stringTodolistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	taskResponse, err := uc.TaskService.CreateTask(todolistID, userID, taskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Task created",
		"body":    taskResponse,
	})
}

func (uc *TaskCont) GetTasks(c *gin.Context) {
	// get payload from token
	userID, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// Get query param
	stringTodolistID := c.Query("todolistID")
	todolistID, err := strconv.ParseInt(stringTodolistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	taskResponse, err := uc.TaskService.GetTasks(todolistID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Task retrieved",
		"body":    taskResponse,
	})
}

func (uc *TaskCont) EditTask(c *gin.Context) {
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
	var taskRequest web.TaskRequest
	if err := c.BindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	// Get query param
	stringTaskID := c.Query("taskID")
	taskID, err := strconv.ParseInt(stringTaskID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}
	stringTodolistID := c.Query("todolistID")
	todolistID, err := strconv.ParseInt(stringTodolistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	taskResponse, err := uc.TaskService.EditTask(taskID, todolistID, userID, taskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Task updated",
		"body":    taskResponse,
	})
}

func (uc *TaskCont) DeleteTask(c *gin.Context) {
	// get payload from token
	userID, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// Get query param
	stringTaskID := c.Query("taskID")
	taskID, err := strconv.ParseInt(stringTaskID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}
	stringTodolistID := c.Query("todolistID")
	todolistID, err := strconv.ParseInt(stringTodolistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	if err := uc.TaskService.DeleteTask(taskID, todolistID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Task deleted",
	})
}
