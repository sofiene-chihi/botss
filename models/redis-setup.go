package models

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {
	fmt.Println("Connecting to Redis")

	addr := "localhost:4000" // Default address

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client

}
