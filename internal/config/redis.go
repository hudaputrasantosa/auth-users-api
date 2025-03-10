package config

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// Global Redis Client
var RedisClient *redis.Client

func InitRedis() {
	connection, _ := ConnectionURLBuilder("redis")
	redisPassword := Config("REDIS_PASSWORD", "")
	db, _ := strconv.Atoi(Config("REDIS_DB", "0"))

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: redisPassword,
		DB:       db,
		PoolSize: 10,
	})

	// Check redis connection
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect Redis:%v", err)
	} else {
		fmt.Println("Redis connected!")
	}

}
