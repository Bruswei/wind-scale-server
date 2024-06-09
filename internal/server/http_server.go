package server

import (
	"fmt"
	"net/http"
	handlers "wind-scale-server/internal/controller/http"
	"wind-scale-server/internal/windspeed/service"
)

type Server interface {
	Start() error
}

type HTTPServer struct {
	Port             string
	windSpeedService service.WindSpeedGetter
}

func (h *HTTPServer) Start() error {
	handler := &handlers.HTTPHandler{WindSpeedService: h.windSpeedService}
	http.HandleFunc("/load", handler.LoadHandler)
	fmt.Println("Server is running on port ", h.Port)

	return http.ListenAndServe(":"+h.Port, nil)
}

func NewServer(port string, windSpeedService *service.WindSpeedService) *HTTPServer {
	return &HTTPServer{
		Port:             port,
		windSpeedService: windSpeedService,
	}
}
