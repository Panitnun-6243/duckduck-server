package config

import "os"

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

func LoadConfig() *Config {
	return &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb+srv://default_uri"),
		DatabaseName:  getEnv("DATABASE_NAME", "default_db_name"),
		JWTSecret:     getEnv("JWT_SECRET_KEY", "default_secret"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":5050"),
		MqttBroker:    getEnv("MQTT_BROKER", "default_broker"),
		MqttClientID:  getEnv("MQTT_CLIENT_ID", "default_client_id"),
		MqttUsername:  getEnv("MQTT_USERNAME", "default_username"),
		MqttPassword:  getEnv("MQTT_PASSWORD", "default_password"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
