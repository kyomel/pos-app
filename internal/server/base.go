package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kyomel/pos-app/internal/config"
	"github.com/kyomel/pos-app/internal/infra/database"
)

func Start() error {
	cfg := config.GetConfig()

	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		return err
	}

	_ = db

	router := chi.NewRouter()

	slog.Info("server "+cfg.App.Name, slog.String("port", cfg.App.Port))
	http.ListenAndServe(":"+cfg.App.Port, router)

	return nil
}
