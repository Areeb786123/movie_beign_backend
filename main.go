package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Get the MongoDB URI from the environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dt := os.Getenv("MONGODB_URI")
	if dt == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	fmt.Println(dt)
}
