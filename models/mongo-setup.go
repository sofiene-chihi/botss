package models

import (
	"context"
	"fmt"
	"os"

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

func InsertConversationMessage(client *mongo.Client) {

	// take the list of messages from the redis on conversation end
	// and set it in mongodb based on id
}
