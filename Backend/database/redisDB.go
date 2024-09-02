package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var Ctx = context.Background()

// InitializeRedisClient initializes the Redis client
func InitializeRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check for the error
	_, err := redisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Redis is not connected: %v", err)
	}
}

// GetRedisClient returns the Redis client, initializing it if necessary
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitializeRedisClient()
	}
	return redisClient
}
