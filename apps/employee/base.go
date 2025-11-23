package employee

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/kyomel/pos-app/apps/employee/create"
)

func InitModule(router *chi.Mux, db *sql.DB) {
	slog.Debug("starting module employee")

	create.Run(router, db)
}
