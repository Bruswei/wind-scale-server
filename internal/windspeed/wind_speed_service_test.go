package windspeed_test

import (
	"context"
	"fmt"
	"testing"
	"time"
	"wind-scale-server/internal/windspeed"

	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (m *MockClient) FetchCurrentWindSpeedData(ctx context.Context, lat, lon float64) (windspeed.WindSpeedRecord, error) {
	return windspeed.WindSpeedRecord{
		Location:  fmt.Sprintf("%f, %f", lat, lon),
		WindSpeed: 5.5,
		Time:      time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
	}, nil
}

type MockDataStorer struct {
	StoredRecords []windspeed.WindSpeedRecord
	Error         error
}

func (m *MockDataStorer) Create(record windspeed.WindSpeedRecord) error {
	m.StoredRecords = append(m.StoredRecords, record)
	return m.Error
}
func (m *MockDataStorer) Read() ([]windspeed.WindSpeedRecord, error) {
	return m.StoredRecords, nil
}

func (m *MockDataStorer) Update(record windspeed.WindSpeedRecord) error {
	return nil
}

func (m *MockDataStorer) Delete(record windspeed.WindSpeedRecord) error {
	return nil
}
func TestFetchAndStoreWindSpeedData(t *testing.T) {
	mockClient := &MockClient{}
	mockDataStorer := &MockDataStorer{}
	windSpeedService := &windspeed.WindSpeedService{
		APIClient: mockClient,
		DataStore: mockDataStorer,
	}

	ctx := context.Background()
	lat := 10.0
	lon := 20.0

	record, err := windSpeedService.FetchAndStoreWindSpeedData(ctx, lat, lon)
	assert.NoError(t, err)
	assert.Equal(t, "10.000000, 20.000000", record.Location)
	assert.Contains(t, mockDataStorer.StoredRecords, record)
}

func TestFetchAndStoreWindSpeedData_InvalidLat(t *testing.T) {
	mockClient := &MockClient{}
	mockDataStorer := &MockDataStorer{}
	windSpeedService := &windspeed.WindSpeedService{
		APIClient: mockClient,
		DataStore: mockDataStorer,
	}

	ctx := context.Background()
	lat := 100.0
	lon := 20.0

	_, err := windSpeedService.FetchAndStoreWindSpeedData(ctx, lat, lon)
	assert.Error(t, err)
	assert.Equal(t, "invalid latitude or longitude", err.Error())
}

func TestFetchAndStoreWindSpeedData_InvalidLong(t *testing.T) {
	mockClient := &MockClient{}
	mockDataStorer := &MockDataStorer{}
	windSpeedService := &windspeed.WindSpeedService{
		APIClient: mockClient,
		DataStore: mockDataStorer,
	}

	ctx := context.Background()
	lat := 20.0
	lon := -220.0

	_, err := windSpeedService.FetchAndStoreWindSpeedData(ctx, lat, lon)
	assert.Error(t, err)
	assert.Equal(t, "invalid latitude or longitude", err.Error())
}

func TestStoreWindSpeedData(t *testing.T) {
	mockDataStorer := &MockDataStorer{}
	windSpeedService := &windspeed.WindSpeedService{
		DataStore: mockDataStorer,
	}

	record := windspeed.WindSpeedRecord{
		Location:  "10.000000, 20.000000",
		Time:      time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
		WindSpeed: 5.5,
	}

	err := windSpeedService.StoreWindSpeedData(record)
	assert.NoError(t, err)
	assert.Contains(t, mockDataStorer.StoredRecords, record)
}
