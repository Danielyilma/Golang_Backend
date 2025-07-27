package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Danielyilma/routes"
)


/**
 * main - Entry point of the application
 *
 * Description:
 * - Loads environment variables from a .env file
 * - Sets up the Gin web server
 * - Registers routes
 * - Starts the server on the specified port
 */

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8090"
	}

	router := gin.Default()

	routes.BasicRoutes(router)

	router.Run(":" + port)
}