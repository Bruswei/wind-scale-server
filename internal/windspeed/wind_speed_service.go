package windspeed

import (
	"context"
)

type WindSpeedService struct {
	APIClient Client
	DataStore DataStorer
}

func (s *WindSpeedService) FetchWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error) {
	wsr, err := s.APIClient.FetchData(ctx, lat, lon)
	if err != nil {
		return WindSpeedRecord{}, err
	}

	return wsr, nil
}

func (s *WindSpeedService) StoreWindSpeedData(record WindSpeedRecord) error {
	return s.DataStore.StoreData(record)
}
