package gateway

import (
	"github.com/RecommenderGo/internal/usecase"
)

type SignupRequest struct {
	Email    string
	Password string
}

type SignupResponse struct {
	Success bool
}

type AuthGateway struct {
	authUsecase usecase.AuthUsecase
}
