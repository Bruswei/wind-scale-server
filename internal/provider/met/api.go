package met

import "context"

type Client interface {
	FetchData(ctx context.Context, lat, lon float64) (interface{}, error)
}
