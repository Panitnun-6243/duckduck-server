package db

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var db *mongo.Database

func Connect() {
	cfg := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Create a context with timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	db = client.Database(cfg.DatabaseName)
	log.Println("Connected to MongoDB!")
}

func GetDB() *mongo.Database {
	return db
}
