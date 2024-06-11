package server

import (
	"fmt"
	"net/http"
	"wind-scale-server/internal/controller"
	"wind-scale-server/internal/windspeed"
)

type Server interface {
	Start() error
}

type HTTPServer struct {
	Port             string
	windSpeedService windspeed.WindSpeedServiceInterface
}

func (h *HTTPServer) Start() error {

	var ctrl controller.Controller = &controller.HTTPController{
		WindSpeedService: h.windSpeedService,
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

func NewServer(port string, windSpeedService windspeed.WindSpeedServiceInterface) *HTTPServer {
	return &HTTPServer{
		Port:             port,
		windSpeedService: windSpeedService,
	}
}
