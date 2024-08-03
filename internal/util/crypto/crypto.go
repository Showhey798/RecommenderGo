//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./crypto.go -destination=./mock/mock.go

package crypto

import "golang.org/x/crypto/bcrypt"

type Crypto interface {
	EncryptPassword(password string) (string, error)
	CompareHashAndPassword(hash, password string) error
}

type CryptoImpl struct{}

func New() *CryptoImpl {
	return &CryptoImpl{}
}

func (c *CryptoImpl) EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (c *CryptoImpl) CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
