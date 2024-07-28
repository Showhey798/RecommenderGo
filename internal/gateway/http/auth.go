package gateway

import (
	"github.com/Showhey798/RecommenderGo/internal/usecase"
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
