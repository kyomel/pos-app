package encryption

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password string) (hash string, err error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashByte), nil
}

func ValidatePassword(hash, password string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
