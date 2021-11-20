package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func LoadEnv(envKey string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env key:", envKey)
	}
	return os.Getenv(envKey)
}
