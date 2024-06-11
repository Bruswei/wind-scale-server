package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"wind-scale-server/internal/weatherservice"
)

type HTTPRequest struct {
	*http.Request
}

func (r *HTTPRequest) Context() context.Context {
	return r.Request.Context()
}

func (r *HTTPRequest) Param(name string) string {
	return r.URL.Query().Get(name)
}

func (r *HTTPRequest) Method() string {
	return r.Request.Method
}

type HTTPResponse struct {
	http.ResponseWriter
}

func (r *HTTPResponse) SetHeader(name, value string) {
	r.ResponseWriter.Header().Set(name, value)
}

func (r *HTTPResponse) Write(statusCode int, body interface{}) error {
	r.ResponseWriter.WriteHeader(statusCode)
	return json.NewEncoder(r.ResponseWriter).Encode(body)
}

type HTTPController struct {
	WeatherService weatherservice.WeatherServiceInterface
}

func (h *HTTPController) HandleWindSpeedLoad(req Request, res Response) error {
	if req.Method() != http.MethodPost {
		res.Write(http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return nil
	}

	latPara := req.Param("lat")
	lngPara := req.Param("lon")

	if latPara == "" || lngPara == "" {
		res.Write(http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return nil
	}

	lat, err := strconv.ParseFloat(latPara, 64)
	if err != nil || lat < -90 || lat > 90 {
		res.Write(http.StatusBadRequest, map[string]string{"error": "Invalid latitude"})
		return nil
	}

	lon, err := strconv.ParseFloat(lngPara, 64)
	if err != nil || lon < -180 || lon > 180 {
		res.Write(http.StatusBadRequest, map[string]string{"error": "Invalid longitude"})
		return nil
	}

	ctx := req.Context()

	_, err = h.WeatherService.FetchAndStoreWindSpeedData(ctx, lat, lon)
	if err != nil {
		res.Write(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch and store wind speed data"})
		return nil
	}

	res.SetHeader("Content-type", "application/json")
	res.Write(http.StatusOK, map[string]string{"message": "ok"})

	return nil
}
