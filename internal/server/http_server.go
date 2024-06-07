package server

import (
	"fmt"
	"net/http"
	handlers "wind-scale-server/internal/handlers/http"
)

type HTTPServer struct {
	Port string
}

func (h *HTTPServer) Start() error {
	http.HandleFunc("/load", handlers.LoadHandler)
	fmt.Println("Server is running on port ", h.Port)

	return http.ListenAndServe(":"+h.Port, nil)
}
