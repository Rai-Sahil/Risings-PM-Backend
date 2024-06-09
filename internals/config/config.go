package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PostgresConnectionString string
	RedisConnectionString    string
	AzureClientID            string
	AzureClientSecretKey     string
	JWTSecretKey             string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PostgresConnectionString: os.Getenv("POSTGRES_CONNECTION_STRING"),
		RedisConnectionString:    os.Getenv("REDIS_CONNECTION_STRING"),
		AzureClientID:            os.Getenv("AZURE_CLIENT_ID"),
		AzureClientSecretKey:     os.Getenv("AZURE_CLIENT_SECRET"),
		JWTSecretKey:             os.Getenv("JWT_SECRET"),
	}
}
