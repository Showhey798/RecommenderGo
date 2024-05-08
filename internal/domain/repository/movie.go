package repository

import (
	"context"

	"recommender.package/pkg/domain/entity"
)

type MovieRepository interface {
	Get(ctx context.Context, id string) (entity.Movie, error)
	GetByIds(ctx context.Context, ids []string) ([]entity.Movie, error)
	Create(ctx context.Context, movie entity.Movie) error
}
