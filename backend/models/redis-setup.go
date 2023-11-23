package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func ConnectRedis() {
	fmt.Println("Connecting to Redis")

	addr := fmt.Sprintf("%s:4000", os.Getenv("REDIS_HOST"))

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
}

func AddMessagesToConversation(conversationId string, messages []MessageItem) {
	for _, obj := range messages {
		serializedObj, err := json.Marshal(obj)
		if err != nil {
			fmt.Println("Error marshalling object:", err)
			return
		}
		err = RedisClient.RPush(conversationId, serializedObj).Err()
		if err != nil {
			fmt.Println("Error pushing object to Redis:", err)
			return
		}
	}
}

func GetConversationById(conversationId string) ([]MessageItem, error) {

	retrievedObjects, err := RedisClient.LRange(conversationId, 0, -1).Result()
	if err != nil {
		fmt.Println("Error retrieving objects from Redis:", err)
		return nil, err
	}

	// Deserialize objects from JSON
	var deserializedMessages []MessageItem
	for _, obj := range retrievedObjects {
		var message MessageItem
		err := json.Unmarshal([]byte(obj), &message)
		if err != nil {
			fmt.Println("Error unmarshalling object:", err)
			return nil, err
		}
		deserializedMessages = append(deserializedMessages, message)
	}

	return deserializedMessages, nil
}
