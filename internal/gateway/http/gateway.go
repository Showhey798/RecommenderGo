package gateway

import (
	"github.com/Showhey798/RecommenderGo/internal/usecase"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Gateway struct {
	logger  *zap.Logger
	usecase usecase.Usecase
}

func NewGateway(logger *zap.Logger, usecase usecase.Usecase) *Gateway {
	return &Gateway{logger, usecase}
}

func (g *Gateway) RegisterGateway(route chi.Router) {
	route.Get("/health", g.HealthCheck)
	route.Post("/signup", g.Signup)
	route.Post("/login", g.Login)
}
