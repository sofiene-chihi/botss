package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func InsertConversationMessage(conversationId string) {

	// take the list of messages from the redis on conversation end
	// and set it in mongodb based on id
	conversationMessages, err := GetConversationById(conversationId)
	if err != nil {
		log.Fatal(err)
	}

	udpatedConversation := Conversation{Messages: conversationMessages}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := MongoClient.Database("chatbot-conversations").Collection("conversation")
	conversationID, err := primitive.ObjectIDFromHex(conversationId)
	if err != nil {
		log.Fatal(err)
	}

	update := bson.M{"$set": udpatedConversation}

	insertResult, err := collection.UpdateByID(ctx, conversationID, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.ModifiedCount)
}
