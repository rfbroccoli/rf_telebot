package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found.")
		return
	}
}

func GetToken() (string, error) {
	token := os.Getenv("TOKEN")
	// log.Printf("token loaded: %v\n", token)
	return token, nil
}

func GetMongoDBClient() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	// log.Printf("uri loaded: %v\n", uri)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, clientOptions)
}
