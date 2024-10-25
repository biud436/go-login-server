package modules

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	// RedisClient 선언
	RedisClient *redis.Client

	// RedisContext 선언
	RedisContext = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: getEnv("REDIS_HOST", "localhost") + ":" + getEnv("REDIS_PORT", "6379"),
	})

	_, err := RedisClient.Ping(RedisContext).Result()
	if err != nil {
		log.Fatalf(("Failed to connect to Redis: %v"), err)
	}

	log.Println("Connected to Redis")
}

func CloseRedis() {
	err := RedisClient.Close()
	if err != nil {
		log.Fatalf("Failed to close Redis connection: %v", err)
	}
}

func IsRedisAvailable() bool {
	_, err := RedisClient.Ping(RedisContext).Result()
	if err != nil {
		return false
	}
	return true
}