package repository

import (
	"context"
	"wind-scale-server/internal/windspeed"
)

// Repository defines the interface for storing data
type Repository interface {
	StoreData(ctx context.Context, data []windspeed.WindSpeedRecord) error
}
