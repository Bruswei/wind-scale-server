package main

import (
	"log"
	"wind-scale-server/internal/config"
	"wind-scale-server/internal/provider/met"
	"wind-scale-server/internal/server"
	"wind-scale-server/internal/windspeed"
)

func main() {
	config := config.GetConfig()

	protocol := "http"

	switch protocol {
	case "gRPC":
		panic("gRPC not implemented")
	default:
		initiateAndRunHTTP(config.ListenPort)
	}

}

func initiateAndRunHTTP(port string) {

	var APIClient windspeed.Client = &met.ExternalClient{}
	var windSpeedService windspeed.WindSpeedGetter = &windspeed.WindSpeedService{
		APIClient: APIClient,
	}

	var srv server.Server = server.NewServer(port, windSpeedService)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
