package token

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	secret := "helo"
	claims := Claims{
		Id:         "1",
		Role:       "admin",
		ExpireTime: time.Duration(time.Now().Add(1 * time.Minute).Unix()),
	}

	token, err := GenerateJWT(claims, secret)
	log.Println(token)

	require.Nil(t, err)
	require.NotEmpty(t, token)
}
