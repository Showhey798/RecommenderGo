package repository

import (
	"context"

	"recommender.package/internal/domain/entity"
)

type MovieRepository interface {
	Get(ctx context.Context, id uint32) (entity.Movie, error)
	GetByIds(ctx context.Context, ids []uint32) ([]entity.Movie, error)
	Create(ctx context.Context, item entity.Movie) error
}
