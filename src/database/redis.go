package database

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RedisClient *redis.Client

func SetupRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("addr"),
		Password: os.Getenv("password"),
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		log.Fatal("Error connecting to redis")
	} else {
		log.Println("Connection with redis successfull")
	}
}
