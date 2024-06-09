package service

import (
	"context"
	"fmt"
	"wind-scale-server/internal/provider/met"
)

type WeatherDataService struct {
	APIClient met.Client
	DPService met.DataProcessingService
}

func (s *WeatherDataService) ProcessData(ctx context.Context, lat, lon float64) (interface{}, error) {
	rawData, err := s.APIClient.FetchData(ctx, lat, lon)
	if err != nil {
		return nil, err
	}

	apiResponse, ok := rawData.(met.APIResponse)
	if !ok {
		return nil, fmt.Errorf("invalid data format")
	}

	processedData, err := s.DPService.ProcessData(apiResponse, lat, lon)
	if err != nil {
		return nil, err
	}

	return processedData, nil
}
