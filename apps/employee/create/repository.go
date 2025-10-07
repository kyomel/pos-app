package create

import "database/sql"

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{db: db}
}
