package main

import (
	"URLShortener/props"
	"URLShortener/server"
	"log"
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
	servers.ConfigureAPI()
}
