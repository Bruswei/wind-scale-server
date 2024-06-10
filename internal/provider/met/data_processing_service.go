package met

import (
	"fmt"
	models "wind-scale-server/internal/windspeed"
)

// WindSpeedData represents the data to be stored in the CSV file
type DataProcessingService interface {
	ProcessData(apiResponse APIResponse, lat, lon float64) ([]models.WindSpeedRecord, error)
	// StoreData(ctx context.Context, data []models.WindSpeedRecord) error
}

type DataService struct {
	// Repository Repository
}

func (s *DataService) ProcessData(apiResponse APIResponse, lat, lon float64) ([]models.WindSpeedRecord, error) {
	var processedData []models.WindSpeedRecord
	location := fmt.Sprintf("%f, %f", lat, lon)

	for _, timeseries := range apiResponse.Properties.Timeseries {
		data := models.WindSpeedRecord{
			Location:  location,
			Time:      timeseries.Time,
			WindSpeed: timeseries.Data.Instant.Details.WindSpeed,
		}
		processedData = append(processedData, data)
	}
	return processedData, nil
}

// Storing data should be moved to windspeed service

// func (s *DataService) StoreData(ctx context.Context, data []models.WindSpeedRecord) error {
// 	// return s.Repository.StoreData(ctx, data)
// 	return nil
// }
