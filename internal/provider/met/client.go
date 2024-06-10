package met

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"wind-scale-server/internal/windspeed"
)

const BaseURL = "https://api.met.no/weatherapi/locationforecast/2.0/compact"

type ExternalClient struct{}

func (c *ExternalClient) FetchData(ctx context.Context, lat, lng float64) ([]windspeed.WindSpeedRecord, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f", BaseURL, lat, lng)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; WindScaleApp/1.0)")
	req.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", response.Status)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(response.Body).Decode((&apiResponse)); err != nil {
		return nil, err
	}

	return c.ProcessData(apiResponse, lat, lng)
}

func (c *ExternalClient) ProcessData(apiResponse APIResponse, lat, lon float64) ([]windspeed.WindSpeedRecord, error) {
	var processedData []windspeed.WindSpeedRecord
	location := fmt.Sprintf("%f, %f", lat, lon)

	for _, timeseries := range apiResponse.Properties.Timeseries {
		data := windspeed.WindSpeedRecord{
			Location:  location,
			Time:      timeseries.Time,
			WindSpeed: timeseries.Data.Instant.Details.WindSpeed,
		}
		processedData = append(processedData, data)
	}
	return processedData, nil
}
