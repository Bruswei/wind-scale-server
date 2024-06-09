package handlers

import "context"

type controller interface {
	LoadCoordinates(ctx context.Context, lat, lon float64) (interface{}, error)
}
