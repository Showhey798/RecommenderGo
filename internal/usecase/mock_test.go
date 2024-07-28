package usecase

import (
	"github.com/Showhey798/RecommenderGo/internal/repository"
	"go.uber.org/mock/gomock"
	"testing"

	mock_repository "github.com/Showhey798/RecommenderGo/internal/repository/mock"
)

type mocks struct {
	auth *mock_repository.MockAuth
}

func newMocks(t *testing.T) *mocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	return &mocks{
		auth: mock_repository.NewMockAuth(ctrl),
	}
}

func newUsecase(m *mocks) *UsecaseImpl {
	return &UsecaseImpl{
		Database: repository.Database{
			Auth: m.auth,
		},
	}
}
