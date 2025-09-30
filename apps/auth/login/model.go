package login

import (
	"time"

	"github.com/kyomel/pos-app/internal/utils/encryption"
	token "github.com/kyomel/pos-app/internal/utils/jwt"
)

type Auth struct {
	ID        int
	PublicId  string
	Email     string
	Password  string
	IsActive  bool
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Auth) ValidatePassword(password string) error {
	if err := encryption.ValidatePassword(a.Password, password); err != nil {
		return errEmailOrPasswordIsNotMatch
	}

	return nil
}

func (a Auth) GenerateToken(data token.Claims, secretKey string) (string, error) {
	token, err := token.GenerateJWT(data, secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
