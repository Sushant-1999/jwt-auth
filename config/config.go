package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// func ConnectDB() {
// 	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017") // Change URI if necessary
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatalf("Failed to ping MongoDB: %v", err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	Client = client
// }

func GetDBCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("DB is not initialized")
	}
	return Client.Database("jwt-auth-db").Collection(collectionName)
}

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading  the .env file")
	}

	MongoDb := os.Getenv("MONGOURI")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB Successfully!")

	Client = client

}
