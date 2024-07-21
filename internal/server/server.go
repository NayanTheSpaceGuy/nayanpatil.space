package server

import (
    "fmt"
	"log"
    "os"

    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
)

func Run() error {

    err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	e := echo.New()

    SetupRoutes(e)

    port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

    fmt.Printf("Starting server on port %s\n", port)
    return e.Start(":" + port)

}
