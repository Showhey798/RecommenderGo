package usecase

import (
	"github.com/Showhey798/RecommenderGo/internal/repository"
)

type Usecase interface {
	// auth usecase
	Signup(params *SignUpParams) error
	Login(params *LoginParams) error
}

type UsecaseImpl struct {
	Database repository.Database
}
