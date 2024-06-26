package main

import (
	"log"
	"wind-scale-server/internal/config"
	"wind-scale-server/internal/csvdata"
	"wind-scale-server/internal/provider/met"
	"wind-scale-server/internal/server"
	"wind-scale-server/internal/weatherservice"
)

func main() {
	config := config.GetConfig()

	// Ensure CSV file exists
	if err := csvdata.FileExists(config.CSVFilePath); err != nil {
		log.Fatalf("Failed to ensure CSV file exists: %v", err)
	}

	protocol := "http"

	switch protocol {
	case "gRPC":
		panic("gRPC not implemented")
	default:
		initiateAndRunHTTP(config.ListenPort, config.CSVFilePath)
	}

}

func initiateAndRunHTTP(port string, filePath string) {

	var APIClient weatherservice.Client = &met.ExternalClient{}
	var CSVStore weatherservice.DataStorer = csvdata.NewCSVStore(filePath)
	var windSpeedService weatherservice.WeatherServiceInterface = &weatherservice.WeatherService{
		APIClient: APIClient,
		DataStore: CSVStore,
	}

	var srv server.Server = server.NewServer(port, windSpeedService)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
