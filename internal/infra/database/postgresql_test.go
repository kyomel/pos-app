package database

import (
	"testing"

	"github.com/kyomel/pos-app/internal/config"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func init() {
	err := config.LoadConfig("../../../config.test.yml")
	if err != nil {
		panic(err)
	}
}

func TestConnectPostgres(t *testing.T) {
	cfg := config.GetConfig()

	db, err := ConnectPostgres(cfg.DB)
	require.Nil(t, err)
	require.NotNil(t, db)
}
