package main

import (
	"log"

	"github.com/joho/godotenv"
	model "github.com/prayogatriady/todolist-app/model/entities"
	"github.com/prayogatriady/todolist-app/repository"
	"github.com/prayogatriady/todolist-app/utils"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Panic(err)
	}

	userRepo := repository.NewUserRepo(db)

	user := model.User{
		Username: "dobow2",
		Password: "pass2",
	}

	log.Println(user)
	// panic("test")

	user, err = userRepo.CreateUser(user)
	if err != nil {
		log.Panic(err)
	}
	log.Println(user)
}
