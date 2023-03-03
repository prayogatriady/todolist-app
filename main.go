package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prayogatriady/todolist-app/controller"
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

	r := gin.Default()

	r.POST("/signup", userCont.Signup)
	r.POST("/signin", userCont.Signin)

	if err := r.Run(":" + PORT); err != nil {
		log.Println(err)
	}
}
