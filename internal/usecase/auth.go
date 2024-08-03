package usecase

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
	"github.com/Showhey798/RecommenderGo/internal/util/errcode"
	"go.uber.org/zap"
)

type SignUpParams struct {
	Email    string
	Password string
}

type LogInParams struct {
	Email    string
	Password string
}

func (u *UsecaseImpl) Signup(ctx context.Context, params *SignUpParams) error {

	password, err := u.Crypto.EncryptPassword(params.Password)

	if err != nil {
		u.logger.Error("failed to encrypt password", zap.Error(err))
		return err
	}

	id := u.IDGenerator.GenerateID()
	user, err := entity.NewUser(id, params.Email, password)

	if err != nil {
		u.logger.Error("failed to create user", zap.Error(err))
		return err
	}

	err = u.Database.Auth.CreateUser(
		ctx,
		user,
	)

	if err != nil {
		u.logger.Error("failed to create user", zap.Error(err))
		return err
	}
	u.logger.Info("user created", zap.String("email", params.Email))
	return nil
}

func (u *UsecaseImpl) Login(ctx context.Context, params *LogInParams) error {
	user, err := u.Database.Auth.FindUserByEmail(ctx, params.Email)
	if err != nil {
		u.logger.Error("failed to find user", zap.Error(err))
		return err
	}
	if err := u.Crypto.CompareHashAndPassword(user.Password, params.Password); err != nil {
		u.logger.Error("password is incorrect")
		return errcode.NewInvalidArgument("password is incorrect")
	}
	return nil

}
