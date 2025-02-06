package config

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// Global Redis Client
var redisClient *redis.Client

func InitRedis() {
	connection, _ := ConnectionURLBuilder("redis")
	redisPassword := Config("REDIS_PASSWORD", "")
	db, _ := strconv.Atoi(Config("REDIS_DB", "0"))

	redisClient = redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: redisPassword,
		DB:       db,
	})

	// Check redis connection
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect Redis:%v", err)
	} else {
		fmt.Println("✅ Redis connected!")
	}

}
