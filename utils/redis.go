package utils

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedisClient() *redis.Client {

	// set environtment variable for setup mysql database
	REDIS_HOST := os.Getenv("REDIS_HOST")
	if REDIS_HOST == "" {
		log.Println("Environment variable REDIS_HOST must be set")
	}

	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: REDIS_PASSWORD,
		DB:       0,
	})

	return client
}
