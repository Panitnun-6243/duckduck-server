package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	MongoURI      string
	JWTSecret     string
	ServerAddress string
	DatabaseName  string
	MqttBroker    string
	MqttClientID  string
	MqttUsername  string
	MqttPassword  string
}

func init() {
	// Load environment variables from .env file for local development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func LoadConfig() *Config {
	return &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DatabaseName:  getEnv("DATABASE_NAME", "default_db_name"),
		JWTSecret:     getEnv("JWT_SECRET_KEY", "default_secret"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":5050"),
		MqttBroker:    getEnv("MQTT_BROKER", "tcp://broker.hivemq.com:1883"),
		MqttClientID:  getEnv("MQTT_CLIENT_ID", "default_client_id"),
		MqttUsername:  getEnv("MQTT_USERNAME", "default_username"),
		MqttPassword:  getEnv("MQTT_PASSWORD", "default_password"),
	}
}

// getEnv gets an environment variable. If it's not set, it returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
