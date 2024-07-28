//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./repository.go -destination=./mock/mock.go

package repository

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
)

type Auth interface {
	// auth repository
	FindUserByEmail(ctx context.Context, email entity.Email) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) error
}

type Database struct {
	Auth Auth
}
