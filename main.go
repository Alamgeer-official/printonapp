package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	startServer()
}

func startServer() {

	// definf port where run the backen server
	port := os.Getenv("SERVER_PORT")
	// port := config.Server_port

	// init Gin router
	r := routes.NewRouter()

	//start server
	r.Run(port)

}
