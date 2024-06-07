package server

import "fmt"

type Server interface {
	Start() error
}

func NewServer(protocol string) (Server, error) {
	switch protocol {
	case "http":
		return &HTTPServer{Port: "8080"}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", protocol)
	}
}
