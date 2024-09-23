package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtSecret string

// LoadConfig loads the values from the .env file once.
func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	JwtSecret = os.Getenv("JWT_SECRET")
}

// GetEnv retrieves the value of the specified key from the environment.
func GetEnv(key string) string {
	return os.Getenv(key)
}
