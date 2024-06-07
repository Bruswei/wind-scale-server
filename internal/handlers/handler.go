package handlers

import "context"

type Handler interface {
	LoadCoordinates(ctx context.Context, lat, lon float64) (interface{}, error)
}
