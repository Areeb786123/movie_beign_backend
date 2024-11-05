package main

import (
	"corses/router"
	"fmt"
	"log"
	"net/http"
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
	r := router.Router()
	fmt.Println("Server is getting started...")

	serr := http.ListenAndServe(":4445", r)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", serr)
	}
	fmt.Println("Listening at port 4000 ...")

}
