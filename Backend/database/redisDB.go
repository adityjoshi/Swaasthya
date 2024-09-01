package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

// Initialize the redis client

var redisClient *redis.Client
var ctx = context.Background()

func InitializeRedisClient() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// check for the error
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis is not connected => %v", err)
	}
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitializeRedisClient()
	}
	return redisClient
}
