package gateway

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/gateway/http/request"
	"github.com/Showhey798/RecommenderGo/internal/usecase"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (g *Gateway) Signup(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	req := &AuthRequest{}

	if err := request.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := g.usecase.Signup(
		ctx,
		&usecase.SignUpParams{
			Email:    req.Email,
			Password: req.Password,
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (g *Gateway) Login(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	req := &AuthRequest{}

	if err := request.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := g.usecase.Login(
		ctx,
		&usecase.LogInParams{
			Email:    req.Email,
			Password: req.Password,
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
