package windspeed

import "context"

type Client interface {
	FetchCurrentWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error)
}

type WindSpeedServiceInterface interface {
	FetchWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error)
	StoreWindSpeedData(record WindSpeedRecord) error
	FetchAndStoreWindSpeedData(ctx context.Context, lat, lon float64) (WindSpeedRecord, error)
}

type DataStorer interface {
	StoreData(WindSpeedRecord) error
}
