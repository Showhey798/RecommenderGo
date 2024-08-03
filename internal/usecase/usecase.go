//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./usecase.go -destination=./mock/mock.go

package usecase

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
	"github.com/Showhey798/RecommenderGo/internal/repository"
	"github.com/Showhey798/RecommenderGo/internal/util/crypto"
	"go.uber.org/zap"
)

type Usecase interface {
	// auth usecase
	Signup(ctx context.Context, params *SignUpParams) error
	Login(ctx context.Context, params *LogInParams) error
}

type UsecaseImpl struct {
	logger      *zap.Logger
	Database    repository.Database
	Crypto      crypto.Crypto
	IDGenerator entity.IDGenerator
}

var _ Usecase = (*UsecaseImpl)(nil)

func New(logger *zap.Logger, database repository.Database) *UsecaseImpl {
	criptoInstance := crypto.New()
	userIDGenerator := entity.NewIDGenerator()
	return &UsecaseImpl{logger: logger, Database: database, Crypto: criptoInstance, IDGenerator: userIDGenerator}
}
