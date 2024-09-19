package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/swaggo/files" // Swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "githuh.com/printonapp/docs" // Import the Swagger docs

	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/routes"
	awssdk "githuh.com/printonapp/utils/aws_sdk"
)

// main initializes the environment and starts the server.
func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	repository.InitDbConnectionos()

	// Initialize AWS session
	awssdk.AwsSessionInit()

	// Start the server
	startServer()
}

// startServer sets up the Gin router, configures routes, and starts the server.
func startServer() {
	// Retrieve the port from environment variables
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "4000" // Default port if not specified
	}

	// Initialize the Gin router
	r := routes.NewRouter()

	// Setup Swagger UI route
	// The Swagger UI will be available at /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("Swagger UI is available at http://localhost:%s/swagger/index.html\n", port)

	// Start the server on the specified port
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
