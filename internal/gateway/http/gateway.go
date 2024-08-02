package gateway

import (
	"github.com/Showhey798/RecommenderGo/internal/usecase"
	"github.com/go-chi/chi"
)

type Gateway struct {
	usecase usecase.Usecase
}

func NewGateway(usecase usecase.Usecase) *Gateway {
	return &Gateway{usecase}
}

func (g *Gateway) RegisterGateway(route chi.Router) {
	route.Post("/signup", g.Signup)
}
