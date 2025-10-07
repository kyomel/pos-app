package create

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Mux, db *sql.DB) {
	slog.Debug("run service create employee", slog.Any("path", "/v1/employees"), slog.Any("method", http.MethodPost))

	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Post("/v1/employee", handler.createEmployee)
}
