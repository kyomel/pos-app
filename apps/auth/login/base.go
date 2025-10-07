package login

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	slog.Debug("run service login", slog.Any("path", "/v1/auth/login"), slog.Any("method", http.MethodPost))

	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Post("/v1/auth/login", handler.Login)
}
