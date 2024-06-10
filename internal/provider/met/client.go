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

func (c *ExternalClient) FetchData(ctx context.Context, lat, lng float64) (windspeed.WindSpeedRecord, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f", BaseURL, lat, lng)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return windspeed.WindSpeedRecord{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; WindScaleApp/1.0)")
	req.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return windspeed.WindSpeedRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return windspeed.WindSpeedRecord{}, fmt.Errorf("failed to fetch data: %s", response.Status)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(response.Body).Decode((&apiResponse)); err != nil {
		return windspeed.WindSpeedRecord{}, err
	}

	return c.getInstantWindSpeed(apiResponse, lat, lng)
}

func (c *ExternalClient) getInstantWindSpeed(apiResponse APIResponse, lat, lon float64) (windspeed.WindSpeedRecord, error) {
	if len(apiResponse.Properties.Timeseries) == 0 {
		return windspeed.WindSpeedRecord{}, fmt.Errorf("no timeseries data available")
	}

	location := fmt.Sprintf("%f, %f", lat, lon)
	timeseries := apiResponse.Properties.Timeseries[0]

	record := windspeed.WindSpeedRecord{
		Location:  location,
		Time:      timeseries.Time,
		WindSpeed: timeseries.Data.Instant.Details.WindSpeed,
	}

	return record, nil
}
