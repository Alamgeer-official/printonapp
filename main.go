package main

import (
	"githuh.com/printonapp/config"
	"githuh.com/printonapp/routes"
)

func main() {
	startServer()
}

func startServer() {
	// definf port where run the backen server
	port := config.Server_port

	// init Gin router
	r := routes.NewRouter()

	//start server
	r.Run(port)

}
