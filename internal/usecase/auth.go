package usecase

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
)

type SignUpParams struct {
	Email    entity.Email
	Password entity.Password
}

type LoginParams struct {
	Email    entity.Email
	Password entity.Password
}

func (u *UsecaseImpl) Signup(ctx context.Context, params *SignUpParams) error {

	err := u.Database.Auth.CreateUser(
		ctx,
		entity.NewUser(params.Email, params.Password),
	)

	if err != nil {
		return err
	}

	return nil
}
