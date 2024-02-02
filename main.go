package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// "githuh.com/printonapp/repository"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/routes"
)

func main() {
	//load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// init database
	repository.InitDbConnectionos()

	// init server
	fmt.Printf("server is running......\n")
	startServer()

}

func startServer() {

	// define port where run the backend server
	port := os.Getenv("SERVER_PORT")
	// port := config.Server_port

	// init Gin router
	r := routes.NewRouter()

	//start server
	err := r.Run(port)
	if err != nil {
		return
	}

}
