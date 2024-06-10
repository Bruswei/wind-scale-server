package windspeed

import (
	"context"
)

type WindSpeedService struct {
	APIClient Client
}

func (s *WindSpeedService) ProcessData(ctx context.Context, lat, lon float64) (interface{}, error) {
	processedData, err := s.APIClient.FetchData(ctx, lat, lon)
	if err != nil {
		return nil, err
	}

	return processedData, nil
}
