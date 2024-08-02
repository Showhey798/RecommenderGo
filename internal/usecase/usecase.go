package usecase

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/repository"
)

type Usecase interface {
	// auth usecase
	Signup(ctx context.Context, params *SignUpParams) error
	//Login(params *LoginParams) error
}

type UsecaseImpl struct {
	Database repository.Database
}

func New(database repository.Database) *UsecaseImpl {
	return &UsecaseImpl{Database: database}
}
