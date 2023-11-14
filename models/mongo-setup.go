package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectToMongoDB() (*mongo.Client, error) {

	mongoUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mongoDatabase := os.Getenv("MONGO_INITDB_DATABASE")

	fmt.Println("database variables")
	fmt.Println(mongoUser, mongoPass, mongoDatabase)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   mongoUser,
		Password:   mongoPass,
		AuthSource: "admin", // Replace with the authentication database
	})

	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return MongoClient, nil
}

func updateConversationMessage(client *mongo.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("chatbot-conversations").Collection("conversation")

	conversationMessages := []MessageItem{
		{Role: "system", Content: "hi there, pretend that you're INSAT administrator during our discussion"},
		{Role: "assistant", Content: "hello there"},
		{Role: "user", Content: "hello, what do you know about INSAT?"},
	}
	newConversation := Conversation{Messages: conversationMessages}
	insertResult, err := collection.InsertOne(ctx, newConversation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document:", insertResult.InsertedID)

}
