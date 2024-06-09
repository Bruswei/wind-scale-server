package service

import "context"

type WindSpeedGetter interface {
	ProcessData(ctx context.Context, lat, lon float64) (interface{}, error)
}
