package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"wind-scale-server/internal/windspeed"
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
	WindSpeedService windspeed.WindSpeedServiceInterface
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
	if err != nil {
		res.Write(http.StatusBadRequest, map[string]string{"error": "Missing lat and/or lon parameters"})
		return nil
	}

	lon, err := strconv.ParseFloat(lngPara, 64)
	if err != nil {
		res.Write(http.StatusBadRequest, map[string]string{"error": "Invalid latitude"})
		return nil
	}

	ctx := req.Context()

	data, err := h.WindSpeedService.FetchWindSpeedData(ctx, lat, lon)
	if err != nil {
		res.Write(http.StatusInternalServerError, map[string]string{"error": "Failed to process data"})
		return nil
	}

	err = h.WindSpeedService.StoreWindSpeedData(data)
	if err != nil {
		res.Write(http.StatusInternalServerError, map[string]string{"error": "Failed to store wind speed data"})
		return nil
	}

	res.SetHeader("Content-type", "application/json")
	res.Write(http.StatusOK, map[string]string{"message": "ok"})

	return nil
}
