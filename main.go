package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prayogatriady/todolist-app/controller"
	"github.com/prayogatriady/todolist-app/middleware"
	"github.com/prayogatriady/todolist-app/repository"
	"github.com/prayogatriady/todolist-app/service"
	"github.com/prayogatriady/todolist-app/utils"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("Environment variable PORT must be set")
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Panic(err)
	}

	userRepo := repository.NewUserRepo(db)
	userServ := service.NewUserServ(userRepo)
	userCont := controller.NewUserCont(userServ)

	todolistRepo := repository.NewListRepo(db)
	todolistServ := service.NewListServ(todolistRepo)
	todolistCont := controller.NewListCont(todolistServ)

	taskRepo := repository.NewTaskRepo(db)
	taskServ := service.NewTaskServ(taskRepo)
	taskCont := controller.NewTaskCont(taskServ)

	r := gin.Default()

	r.POST("/signup", userCont.Signup)
	r.POST("/signin", userCont.Signin)

	// Middleware for authentication
	r.Use(middleware.AuthMiddleware)

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
