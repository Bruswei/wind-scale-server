package windspeed

import "context"

type Client interface {
	FetchData(ctx context.Context, lat, lon float64) ([]WindSpeedRecord, error)
}

type WindSpeedGetter interface {
	ProcessData(ctx context.Context, lat, lon float64) (interface{}, error)
type DataStorer interface {
	StoreData(WindSpeedRecord) error
}
