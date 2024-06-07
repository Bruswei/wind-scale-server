package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func LoadHandler(w http.ResponseWriter, r *http.Request) {
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

	_, err := strconv.ParseFloat(latPara, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	_, err = strconv.ParseFloat(lonPara, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("Ok")
}
