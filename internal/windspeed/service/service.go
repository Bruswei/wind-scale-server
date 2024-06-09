package service

import "context"

type CoordinateService interface {
	ProcessData(ctx context.Context, lat, lon float64) (interface{}, error)
}
