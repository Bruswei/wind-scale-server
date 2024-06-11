package windspeed

import (
	"context"
	"fmt"
)

type WindSpeedService struct {
	APIClient Client
	DataStore DataStorer
}

func (s *WindSpeedService) FetchWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error) {
	if !isValidLatLon(lat, lon) {
		return WindSpeedRecord{}, fmt.Errorf("invalid latitude or longitude")
	}
	wsr, err := s.APIClient.FetchData(ctx, lat, lon)
	if err != nil {
		return WindSpeedRecord{}, err
	}

	return wsr, nil
}

func (s *WindSpeedService) StoreWindSpeedData(record WindSpeedRecord) error {
	return s.DataStore.StoreData(record)
}

func isValidLatLon(lat, lon float64) bool {
	return lat >= -90 && lat <= 90 && lon >= -180 && lon <= 180
}
