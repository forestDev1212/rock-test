package database

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"

	"rocks-test/config"
)

var RedisClient *redis.Client

func SetupRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})

	// Ping Redis to ensure the connection is established
	pong, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Printf("Connected to Redis successfully: %s\n", pong)
}
