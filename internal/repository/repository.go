package repository

import (
	"context"
	"wind-scale-server/internal/models"
)

// Repository defines the interface for storing data
type Repository interface {
	StoreData(ctx context.Context, data []models.WindSpeedData) error
}
