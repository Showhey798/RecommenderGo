package usecase

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/util/errcode"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
)

func TestUsecaseImpl_Signup(t *testing.T) {

	userID := "test-user-1"
	validEmail := "test@gmail.com"
	validPassword := "password"
	encryptoPassword := "encryptedPassword"

	user, _ := entity.NewUser(userID, validEmail, encryptoPassword)

	t.Parallel()
	ctx := context.Background()
	tests := []struct {
		name    string
		setup   func(*testing.T, *mocks)
		params  *SignUpParams
		wantErr bool
	}{
		{
			name: "failed to create user already exists",
			setup: func(t *testing.T, m *mocks) {
				m.auth.EXPECT().CreateUser(
					ctx,
					user,
				).Return(errcode.NewAlreadyExists("email: %s", validEmail))
				m.crypto.EXPECT().EncryptPassword(validPassword).Return(encryptoPassword, nil)
				m.idgenerator.EXPECT().GenerateID().Return(userID)
			},
			params: &SignUpParams{
				Email:    validEmail,
				Password: validPassword,
			},
			wantErr: true,
		},
		{
			name: "success to create user",
			setup: func(t *testing.T, m *mocks) {
				m.auth.EXPECT().CreateUser(
					ctx,
					user,
				).Return(nil)
				m.crypto.EXPECT().EncryptPassword(validPassword).Return(encryptoPassword, nil)
				m.idgenerator.EXPECT().GenerateID().Return(userID)
			},
			params: &SignUpParams{
				Email:    validEmail,
				Password: validPassword,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(t, m)
			u := newUsecase(m)
			err := u.Signup(ctx, tt.params)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestUsecaseImpl_Login(t *testing.T) {

	userID := "test-user-1"
	validEmail := "test@gmail.com"
	validPassword := "password"
	HashPassword := "encryptedPassword"

	user := entity.User{
		ID:       userID,
		Email:    validEmail,
		Password: HashPassword,
	}

	t.Parallel()

	ctx := context.Background()
	tests := []struct {
		name    string
		setup   func(*testing.T, *mocks)
		params  *LogInParams
		wantErr bool
	}{
		{
			name: "failed to find user",
			setup: func(t *testing.T, m *mocks) {
				m.auth.EXPECT().FindUserByEmail(
					ctx,
					validEmail,
				).Return(entity.User{}, errcode.NewNotFound("email: %s", validEmail))
			},
			params: &LogInParams{
				Email:    validEmail,
				Password: validPassword,
			},
			wantErr: true,
		},
		{
			name: "password is incorrect",
			setup: func(t *testing.T, m *mocks) {
				m.auth.EXPECT().FindUserByEmail(
					ctx,
					validEmail,
				).Return(user, nil)
				m.crypto.EXPECT().CompareHashAndPassword(HashPassword, validPassword).Return(errcode.NewInvalidArgument("password is incorrect"))
			},
			params: &LogInParams{
				Email:    validEmail,
				Password: validPassword,
			},
			wantErr: true,
		},
		{
			name: "success to login",
			setup: func(t *testing.T, m *mocks) {
				m.auth.EXPECT().FindUserByEmail(
					ctx,
					validEmail,
				).Return(user, nil)
				m.crypto.EXPECT().CompareHashAndPassword(HashPassword, validPassword).Return(nil)
			},
			params: &LogInParams{
				Email:    validEmail,
				Password: validPassword,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(t, m)
			u := newUsecase(m)
			err := u.Login(ctx, tt.params)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
