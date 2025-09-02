package main

import (
	"github.com/kyomel/pos-app/internal/config"
	"github.com/kyomel/pos-app/internal/infra/database"
)

func main() {
	err := config.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		panic(err)
	}

	_ = db
}
