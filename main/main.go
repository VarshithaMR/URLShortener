package main

import (
	"log"

	"URLShortener/props"
	"URLShortener/server"
	"URLShortener/service"
)

func main() {
	startApplication()
}

func startApplication() {
	properties, err := props.ReadProperties("./env/application.yaml")
	if err != nil {
		log.Fatalf("Error reading configurations file: %v", err)
	}

	servers := server.NewServer(properties)

	newHandler := service.NewURLShortener()
	servers.ConfigureAPI(newHandler)
}
