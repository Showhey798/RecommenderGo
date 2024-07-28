package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "success",
			user: User{
				Email:    Email("test@gmail.com"),
				Password: Password("password"),
			},
			want: true,
		},
		{
			name: "fail @ is not included",
			user: User{
				Email:    Email("testgmail.com"),
				Password: Password("password"),
			},
			want: false,
		},
		{
			name: "fail it has more than one @",
			user: User{
				Email:    Email("test@gmail@com"),
				Password: Password("password"),
			},
			want: false,
		},
		{
			name: "fail it has different domain",
			user: User{
				Email:    Email("test@example.com"),
				Password: Password("password"),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.user.Validate()

			require.Equal(t, tt.want, got)
		})
	}
}
