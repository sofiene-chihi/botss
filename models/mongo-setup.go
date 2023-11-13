package models

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {

	mongoUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mongoDatabase := os.Getenv("MONGO_INITDB_DATABASE")

	connectionURI := fmt.Sprintf("mongodb://%s:%s@localhost:27017/%s", mongoUser, mongoPass, mongoDatabase)
	clientOptions := options.Client().ApplyURI(connectionURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
