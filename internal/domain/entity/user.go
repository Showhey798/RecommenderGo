package entity

import "github.com/google/uuid"

type UserID string
type Password string

type User struct {
	ID       UserID
	Email    Email
	Password Password
}

func NewUser(email Email, password Password) *User {
	return &User{
		ID:       UserID(uuid.New().String()),
		Email:    email,
		Password: password,
	}
}

func (u *User) Validate() bool {
	return u.Email.CheckAvailableEmail()
}
