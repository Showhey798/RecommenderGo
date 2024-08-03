//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./user.go -destination=./mock/mock_user.go

package entity

import (
	"github.com/Showhey798/RecommenderGo/internal/util/errcode"
	"github.com/google/uuid"
	"strings"
)

type AvailableEmailDomain string

const (
	GmailDomain AvailableEmailDomain = "gmail.com"
)

type User struct {
	ID       string
	Email    string
	Password string
}

func NewUser(ID string, email string, password string) (*User, error) {
	user := &User{
		ID:       ID,
		Email:    email,
		Password: password,
	}
	if !user.validateEmail() {
		return nil, errcode.NewInvalidArgument("invalid email")
	}
	return user, nil
}

func (u *User) validateEmail() bool {
	// @以降がgmail.comのみ許可
	userDomains := strings.Split(u.Email, "@")
	if len(userDomains) != 2 {
		return false
	}
	if userDomains[1] != string(GmailDomain) {
		return false
	}
	return true
}

type IDGenerator interface {
	GenerateID() string
}
type UserIDGenerator struct{}

func NewIDGenerator() *UserIDGenerator {
	return &UserIDGenerator{}
}

func (u *UserIDGenerator) GenerateID() string {
	return uuid.New().String()
}
