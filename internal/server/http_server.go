package server

import (
	"fmt"
	"net/http"
	"wind-scale-server/internal/api"
	handlers "wind-scale-server/internal/handlers/http"
	"wind-scale-server/internal/service"
)

type HTTPServer struct {
	Port string
}

func (h *HTTPServer) Start() error {
	APIClient := &api.ExternalClient{}
	dPService := &service.DataService{}
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
