package usecase

import (
	"context"
	"github.com/Showhey798/RecommenderGo/internal/util/errcode"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
)

func TestUsecaseImpl_Signup(t *testing.T) {

	validEmail := entity.Email("test@gmail.com")
	validPassword := entity.Password("password")

	t.Parallel()
	ctx := context.Background()
	tests := []struct {
		name    string
		setup   func(*testing.T, *mocks)
		params  *SignUpParams
		want    bool
		wantErr bool
	}{
		{
			name: "failed to create user already exists",
			setup: func(t *testing.T, m *mocks) {
				m.auth.EXPECT().CreateUser(
					ctx,
					entity.User{
						Email:    validEmail,
						Password: validPassword,
					},
				).Return(errcode.NewAlreadyExists("email: %s", validEmail))
			},
			params: &SignUpParams{
				Email:    validEmail,
				Password: validPassword,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "success to create user",
			setup: func(t *testing.T, m *mocks) {
				m.auth.EXPECT().CreateUser(
					ctx,
					entity.User{
						Email:    validEmail,
						Password: validPassword,
					},
				).Return(nil)
			},
			params: &SignUpParams{
				Email:    validEmail,
				Password: validPassword,
			},
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(t, m)
			u := newUsecase(m)
			got, err := u.Signup(ctx, tt.params)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, got)
		})
	}
}
