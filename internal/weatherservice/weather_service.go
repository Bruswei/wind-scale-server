package weatherservice

import (
	"context"
	"fmt"
)

type WeatherService struct {
	APIClient Client
	DataStore DataStorer
}

func (s *WeatherService) FetchWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error) {
	if !isValidLatLon(lat, lon) {
		return WindSpeedRecord{}, fmt.Errorf("invalid latitude or longitude")
	}
	wsr, err := s.APIClient.FetchCurrentWindSpeedData(ctx, lat, lon)
	if err != nil {
		return WindSpeedRecord{}, err
	}

	return wsr, nil
}

func (s *WeatherService) StoreWindSpeedData(record WindSpeedRecord) error {
	return s.DataStore.Create(record)
}

func (s *WeatherService) FetchAndStoreWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error) {
	record, err := s.FetchWindSpeedData(ctx, lat, lon)
	if err != nil {
		return WindSpeedRecord{}, err
	}

	err = s.StoreWindSpeedData(record)
	if err != nil {
		return WindSpeedRecord{}, err
	}

	return record, nil
}

func isValidLatLon(lat, lon float64) bool {
	return lat >= -90 && lat <= 90 && lon >= -180 && lon <= 180
}
