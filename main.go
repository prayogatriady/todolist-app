package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/todolist-app/controller"
	"github.com/prayogatriady/todolist-app/middleware"
	"github.com/prayogatriady/todolist-app/repository"
	"github.com/prayogatriady/todolist-app/service"
	"github.com/prayogatriady/todolist-app/utils"
)

func main() {
	// set environtment variable for for PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("Environment variable PORT must be set")
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Panic(err)
	}

	// access each layer for User
	userRepo := repository.NewUserRepo(db)
	userServ := service.NewUserServ(userRepo)
	userCont := controller.NewUserCont(userServ)

	// access each layer for Todolist
	todolistRepo := repository.NewListRepo(db)
	todolistServ := service.NewListServ(todolistRepo)
	todolistCont := controller.NewListCont(todolistServ)

	// access each layer for Task
	taskRepo := repository.NewTaskRepo(db)
	taskServ := service.NewTaskServ(taskRepo)
	taskCont := controller.NewTaskCont(taskServ)

	// init instance for Gin framework
	r := gin.Default()

	// routes that can be accessed with out authentication
	r.POST("/signup", userCont.Signup)
	r.POST("/signin", userCont.Signin)

	// middleware for authentication
	r.Use(middleware.AuthMiddleware)

	// routes that can be accessed with authentication
	api := r.Group("/api")
	{
		api.GET("/profile", userCont.Profile)
		api.PUT("/edit", userCont.EditProfile)
		api.DELETE("/delete", userCont.DeleteUser)

		api.GET("/todolist", todolistCont.GetList)
		api.POST("/todolist", todolistCont.CreateTodolist)
		api.PUT("/todolist/:todolistID", todolistCont.EditList)
		api.DELETE("/todolist/:todolistID", todolistCont.DeleteList)

		api.GET("/todolist/task", taskCont.GetTasks)
		api.POST("/todolist/task", taskCont.CreateTask)
		api.PUT("/todolist/task", taskCont.EditTask)
		api.DELETE("/todolist/task", taskCont.DeleteTask)
	}

	if err := r.Run(":" + PORT); err != nil {
		log.Println(err)
	}
}
