package usecase

import (
	"github.com/Showhey798/RecommenderGo/internal/repository"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"

	mock_user "github.com/Showhey798/RecommenderGo/internal/domain/entity/mock"
	mock_repository "github.com/Showhey798/RecommenderGo/internal/repository/mock"
	mock_crypto "github.com/Showhey798/RecommenderGo/internal/util/crypto/mock"
)

type mocks struct {
	auth        *mock_repository.MockAuth
	crypto      *mock_crypto.MockCrypto
	idgenerator *mock_user.MockIDGenerator
}

func newMocks(t *testing.T) *mocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	return &mocks{
		auth:        mock_repository.NewMockAuth(ctrl),
		crypto:      mock_crypto.NewMockCrypto(ctrl),
		idgenerator: mock_user.NewMockIDGenerator(ctrl),
	}
}

func newUsecase(m *mocks) *UsecaseImpl {
	logger, _ := zap.NewProduction()
	return &UsecaseImpl{
		logger: logger,
		Database: repository.Database{
			Auth: m.auth,
		},
		Crypto:      m.crypto,
		IDGenerator: m.idgenerator,
	}
}
