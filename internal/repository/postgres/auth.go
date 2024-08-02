package postgres

import (
	"context"
	"database/sql"
	"github.com/Showhey798/RecommenderGo/internal/domain/entity"
	"github.com/Showhey798/RecommenderGo/internal/util/errcode"
)

type AuthRepository struct {
	DB *sql.DB
}

func (a *AuthRepository) CreateUser(ctx context.Context, user *entity.User) error {
	var userID entity.UserID
	err := a.DB.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if userID != "" {
		return errcode.NewAlreadyExists("email: %s", user.Email)
	}

	if _, err := a.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password); err != nil {
		return err
	}
	return nil
}
