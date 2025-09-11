package config

import (
	"Evaluation_Service/config/env"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func CreateRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: env.GetString("REDIS_URL", "localhost:6379"),
		DB:   env.GetInt("REDIS_DB", 0),
		Password: env.GetString("REDIS_PASSWORD", ""),
	})

	_ , err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil
	}
	return client
}