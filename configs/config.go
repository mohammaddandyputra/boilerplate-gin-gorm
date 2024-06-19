package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUsername string
	DBPassword string
	DBName     string
}

func LoadConfig() Config {
	return Config{
		DBHost:     "localhost",
		DBPort:     5432,
		DBUsername: "postgres",
		DBPassword: "postgres",
		DBName:     "learn_gin_gorm",
	}
}

func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
