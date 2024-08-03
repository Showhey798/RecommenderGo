package postgres

import (
	"context"
	"database/sql"
	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
	"github.com/Showhey798/RecommenderGo/internal/repository"
	"github.com/Showhey798/RecommenderGo/internal/util/errcode"
)

var _ repository.Auth = (*AuthRepository)(nil)

type AuthRepository struct {
	DB *sql.DB
}

func (a *AuthRepository) CreateUser(ctx context.Context, user *entity.User) error {
	var userID string
	err := a.DB.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if userID != "" {
		return errcode.NewAlreadyExists("email: %s", user.Email)
	}

	if _, err := a.DB.Exec("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password); err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := a.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
