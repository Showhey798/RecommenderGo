package gateway

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
	"github.com/Showhey798/RecommenderGo/internal/gateway/http/middleware"
	"github.com/Showhey798/RecommenderGo/internal/gateway/http/request"
	"github.com/Showhey798/RecommenderGo/internal/usecase"
	"net/http"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Gateway) Signup(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	req := &SignupRequest{}

	if err := request.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password, err := middleware.EncryptPassword(req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.usecase.Signup(
		ctx,
		&usecase.SignUpParams{
			Email:    entity.Email(req.Email),
			Password: entity.Password(password),
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
