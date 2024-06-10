package windspeed

import (
	"context"
)

// Repository defines the interface for storing data
type DataStorer interface {
	StoreData(ctx context.Context, data []WindSpeedRecord) error
}
