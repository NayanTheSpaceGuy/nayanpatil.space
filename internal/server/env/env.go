package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Port string

func Init() error {
    err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

    Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080" // Default port if not specified
        log.Println("PORT environment variable not set, using default value 8080")
	}

    return nil
}
