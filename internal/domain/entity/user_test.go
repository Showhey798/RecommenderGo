package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser_ValidateEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "success",
			user: User{
				Email:    "test@gmail.com",
				Password: "password",
			},
			want: true,
		},
		{
			name: "fail @ is not included",
			user: User{
				Email:    "testgmail.com",
				Password: "password",
			},
			want: false,
		},
		{
			name: "fail it has more than one @",
			user: User{
				Email:    "test@gmail@com",
				Password: "password",
			},
			want: false,
		},
		{
			name: "fail it has different domain",
			user: User{
				Email:    "test@example.com",
				Password: "password",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.user.ValidateEmail()

			require.Equal(t, tt.want, got)
		})
	}
}
