package server

import (
	"fmt"
	"net/http"
	"wind-scale-server/internal/controller"
	"wind-scale-server/internal/weatherservice"
)

type Server interface {
	Start() error
}

type HTTPServer struct {
	Port           string
	WeatherService weatherservice.WeatherServiceInterface
}

func (h *HTTPServer) Start() error {

	var ctrl controller.Controller = &controller.HTTPController{
		WeatherService: h.WeatherService,
	}

	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		req := &controller.HTTPRequest{Request: r}
		res := &controller.HTTPResponse{ResponseWriter: w}
		err := ctrl.HandleWindSpeedLoad(req, res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server is running on port ", h.Port)
	return http.ListenAndServe(":"+h.Port, nil)
}

func NewServer(port string, ws weatherservice.WeatherServiceInterface) *HTTPServer {
	return &HTTPServer{
		Port:           port,
		WeatherService: ws,
	}
}
