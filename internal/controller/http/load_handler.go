package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wind-scale-server/internal/windspeed/service"
)

type HTTPHandler struct {
	CoordinateService service.CoordinateService
}

func (h *HTTPHandler) LoadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	latPara := r.URL.Query().Get("lat")
	lonPara := r.URL.Query().Get("lon")

	if latPara == "" || lonPara == "" {
		http.Error(w, "Missing lat and/or lon paramters", http.StatusBadRequest)
		return
	}

	lat, err := strconv.ParseFloat(latPara, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonPara, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	data, err := h.CoordinateService.ProcessData(ctx, lat, lon)

	fmt.Println(data)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("Ok")
}
