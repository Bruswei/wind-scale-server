package server

import (
	"fmt"
	"net/http"
	handlers "wind-scale-server/internal/controller/http"
	"wind-scale-server/internal/provider/met"
	"wind-scale-server/internal/windspeed/service"
)

type HTTPServer struct {
	Port string
}

func (h *HTTPServer) Start() error {
	APIClient := &met.ExternalClient{}
	dPService := &met.DataService{}
	windScaleAPIService := &service.WeatherDataService{
		APIClient: APIClient,
		DPService: dPService,
	}
	handler := &handlers.HTTPHandler{CoordinateService: windScaleAPIService}
	// Repository instanilize
	http.HandleFunc("/load", handler.LoadHandler)
	fmt.Println("Server is running on port ", h.Port)

	return http.ListenAndServe(":"+h.Port, nil)
}
