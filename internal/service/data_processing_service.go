package service

import (
	"context"
	"fmt"
	"wind-scale-server/internal/models"
)

// ProcessedData represents the data to be stored in the CSV file
type ProcessedData struct {
	Location  string
	Time      string
	WindSpeed float64
}

type DataProcessingService interface {
	ProcessData(apiResponse models.APIResponse, lat, lon float64) ([]ProcessedData, error)
	StoreData(ctx context.Context, data []ProcessedData) error
}

type DataService struct {
	// Repository Repository
}

func (s *DataService) ProcessData(apiResponse models.APIResponse, lat, lon float64) ([]ProcessedData, error) {
	var processedData []ProcessedData
	location := fmt.Sprintf("%f, %f", lat, lon)

	for _, timeseries := range apiResponse.Properties.Timeseries {
		data := ProcessedData{
			Location:  location,
			Time:      timeseries.Time,
			WindSpeed: timeseries.Data.Instant.Details.WindSpeed,
		}
		processedData = append(processedData, data)
	}
	return processedData, nil
}

func (s *DataService) StoreData(ctx context.Context, data []ProcessedData) error {
	// return s.Repository.StoreData(ctx, data)
	return nil
}
