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

func (u *UsecaseImpl) Signup(ctx context.Context, params *SignUpParams) (bool, error) {
	// 同一のemailが存在するか確認
	err := u.Database.Auth.CreateUser(
		ctx,
		entity.User{
			Email:    params.Email,
			Password: params.Password,
		})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UsecaseImpl) Login(ctx context.Context, params *LoginParams) (bool, error) {
	return false, nil
}
