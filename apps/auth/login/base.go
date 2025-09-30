package login

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	slog.Debug("run service login", slog.Any("path", "/v1/auth/login"))

	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Post("/v1/auth/login", handler.Login)
}
