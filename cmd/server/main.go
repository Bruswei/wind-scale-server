package main

import (
	"log"
	"wind-scale-server/internal/provider/met"
	"wind-scale-server/internal/server"
	"wind-scale-server/internal/windspeed/service"
)

func main() {
	protocol := "http"

	switch protocol {
	case "gRPC":
		panic("gRPC not implemented")
	default:
		port := "8080"
		initiateAndRunHTTP(port)
	}

}

func initiateAndRunHTTP(port string) {

	var APIClient met.Client = &met.ExternalClient{}
	var dPService met.DataProcessingService = &met.DataService{}
	var windSpeedService service.WindSpeedGetter = &service.WindSpeedService{
		APIClient: APIClient,
		DPService: dPService,
	}

	var srv server.Server = server.NewServer(port, windSpeedService)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
