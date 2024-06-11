package met

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wind-scale-server/internal/weatherservice"
)

const BaseURL = "https://api.met.no/weatherapi/locationforecast/2.0/compact"

type ExternalClient struct{}

func (c *ExternalClient) FetchCurrentWindSpeedData(ctx context.Context, lat, lng float64) (weatherservice.WindSpeedRecord, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f", BaseURL, lat, lng)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return weatherservice.WindSpeedRecord{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; WindScaleApp/1.0)")
	req.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return weatherservice.WindSpeedRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return weatherservice.WindSpeedRecord{}, fmt.Errorf("failed to fetch data: %s", response.Status)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(response.Body).Decode((&apiResponse)); err != nil {
		return weatherservice.WindSpeedRecord{}, err
	}

	return c.findClosestTimeSeries(apiResponse, lat, lng)
}

func (c *ExternalClient) findClosestTimeSeries(apiResponse APIResponse, lat, lon float64) (weatherservice.WindSpeedRecord, error) {
	if len(apiResponse.Properties.Timeseries) == 0 {
		return weatherservice.WindSpeedRecord{}, fmt.Errorf("no timeseries data available")
	}

	location := fmt.Sprintf("%f, %f", lat, lon)
	currentTime := time.Now()
	var closestTimeSeries Timeseries
	minDiff := time.Duration(1<<63 - 1)

	for _, ts := range apiResponse.Properties.Timeseries {
		tsTime, err := time.Parse(time.RFC3339, ts.Time)
		if err != nil {
			continue
		}

		diff := currentTime.Sub(tsTime)
		if diff < 0 {
			diff = -diff
		}

		if diff < minDiff {
			minDiff = diff
			closestTimeSeries = ts
		}
	}

	if closestTimeSeries.Time == "" {
		return weatherservice.WindSpeedRecord{}, fmt.Errorf("no valid timeseries data found")
	}

	record := weatherservice.WindSpeedRecord{
		Location:  location,
		Time:      closestTimeSeries.Time,
		WindSpeed: closestTimeSeries.Data.Instant.Details.WindSpeed,
	}

	return record, nil
}
