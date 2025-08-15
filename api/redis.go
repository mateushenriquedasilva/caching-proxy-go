package api

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redis_database *redis.Client
var ctx = context.Background()

func InitRedis() {
	godotenv.Load()
	redis_password := os.Getenv("REDIS_PASSWORD")

	redis_database = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: redis_password,
		DB:       0,
	})

	_, err := redis_database.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis connection successfully established!")
}
