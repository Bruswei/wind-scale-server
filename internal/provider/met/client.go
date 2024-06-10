package met

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://api.met.no/weatherapi/locationforecast/2.0/compact"

type Client interface {
	FetchData(ctx context.Context, lat, lon float64) (interface{}, error)
}

type ExternalClient struct{}

func (c *ExternalClient) FetchData(ctx context.Context, lat, lon float64) (interface{}, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f", BaseURL, lat, lon)
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

	return apiResponse, nil
}
