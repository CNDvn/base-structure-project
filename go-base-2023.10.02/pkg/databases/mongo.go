package databases

import (
	"context"
	"fmt"
	"gobase/pkg/helpers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func (t *MongoClient) Connection() *MongoClient {
	// Set up MongoDB client options
	uri := helpers.GetENV().MONGO_URI
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB!")

	t.Client = client
	return t
}

func (t *MongoClient) Close() {
	if err := t.Client.Disconnect(context.TODO()); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Disconnected from MongoDB!")
}
