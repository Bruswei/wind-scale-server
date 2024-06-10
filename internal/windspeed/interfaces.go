package windspeed

import "context"

type Client interface {
	FetchData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error)
}

type WindSpeedGetter interface {
	FetchWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error)
	StoreWindSpeedData(record WindSpeedRecord) error
}

type DataStorer interface {
	StoreData(WindSpeedRecord) error
}
