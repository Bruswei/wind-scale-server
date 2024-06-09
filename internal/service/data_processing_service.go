package service

import (
	"context"
	"fmt"
	"wind-scale-server/internal/models"
)

// WindSpeedData represents the data to be stored in the CSV file

type DataProcessingService interface {
	ProcessData(apiResponse models.APIResponse, lat, lon float64) ([]models.WindSpeedData, error)
	StoreData(ctx context.Context, data []models.WindSpeedData) error
}

// DataService is here but the interface should be moved out to its own file?

type DataService struct {
	// Repository Repository
}

func (s *DataService) ProcessData(apiResponse models.APIResponse, lat, lon float64) ([]models.WindSpeedData, error) {
	var processedData []models.WindSpeedData
	location := fmt.Sprintf("%f, %f", lat, lon)

	for _, timeseries := range apiResponse.Properties.Timeseries {
		data := models.WindSpeedData{
			Location:  location,
			Time:      timeseries.Time,
			WindSpeed: timeseries.Data.Instant.Details.WindSpeed,
		}
		processedData = append(processedData, data)
	}
	return processedData, nil
}

func (s *DataService) StoreData(ctx context.Context, data []models.WindSpeedData) error {
	// return s.Repository.StoreData(ctx, data)
	return nil
}
