package main

import (
	"log"
	"wind-scale-server/internal/server"
)

func main() {
	protocol := "http"

	srv, err := server.NewServer(protocol)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
