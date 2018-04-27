// main.go

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASS"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_SSLMODE"))
	a.Run(":8080")
}
