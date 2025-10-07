package employee

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func InitModule(router *chi.Mux, db *sql.DB) {
	slog.Debug("starting module employee")
}
