package gateway

import (
	"github.com/Showhey798/RecommenderGo/internal/usecase/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

type mocks struct {
	uc *mock.MockUsecase
}

func newMocks(t *testing.T) *mocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	return &mocks{
		uc: mock.NewMockUsecase(ctrl),
	}
}
