package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	filename := "../../config.test.yml"
	err := LoadConfig(filename)

	require.Nil(t, err)
	require.NotNil(t, cfg)
	require.Equal(t, "localhost", cfg.DB.Host)
	require.Equal(t, "5432", cfg.DB.Port)
}
