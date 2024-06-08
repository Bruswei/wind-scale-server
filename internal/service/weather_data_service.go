package service

import (
	"context"
	"fmt"
	"wind-scale-server/internal/api"
	"wind-scale-server/internal/models"
)

type WeatherDataService struct {
	APIClient api.Client
	DPService DataProcessingService
}

func (s *WeatherDataService) ProcessData(ctx context.Context, lat, lon float64) (interface{}, error) {
	rawData, err := s.APIClient.FetchData(ctx, lat, lon)
	if err != nil {
		return nil, err
	}

	apiResponse, ok := rawData.(models.APIResponse)
	if !ok {
		return nil, fmt.Errorf("invalid data format")
	}

	processedData, err := s.DPService.ProcessData(apiResponse, lat, lon)
	if err != nil {
		return nil, err
	}

	return processedData, nil
}
