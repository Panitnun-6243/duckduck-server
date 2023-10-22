package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
	"time"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

var (
	client *mongo.Client
)

// InitializeDB initializes the database connection
func InitializeDB() {
	loadEnv()

	MongoURI := os.Getenv("MONGO_URI")
	//DatabaseName := os.Getenv("DATABASE_NAME")
	ConnectionTimeout := os.Getenv("CONNECTION_TIMEOUT")

	connectionTimeoutInt, err := strconv.Atoi(ConnectionTimeout)
	if err != nil {
		log.Fatalf("Failed to convert CONNECTION_TIMEOUT to int: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(connectionTimeoutInt)*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB successfully")
}

// GetDB returns a reference to the database
func GetDB() *mongo.Database {
	DatabaseName := os.Getenv("DATABASE_NAME")
	return client.Database(DatabaseName)
}

// DisconnectDB closes the database connection
func DisconnectDB() {
	ConnectionTimeout := os.Getenv("CONNECTION_TIMEOUT")
	connectionTimeoutInt, err := strconv.Atoi(ConnectionTimeout)
	if err != nil {
		log.Fatalf("Failed to convert CONNECTION_TIMEOUT to int: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(connectionTimeoutInt)*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
	log.Println("Connection to MongoDB closed.")
}
