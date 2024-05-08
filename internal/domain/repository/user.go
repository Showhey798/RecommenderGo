package repository

import (
	"context"

	"recommender.package/internal/domain/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (entity.User, error)
	Create(ctx context.Context, user entity.User) error
}
