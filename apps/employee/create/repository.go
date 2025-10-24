package create

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// Begin implements repoContract.
func (r repository) Begin(ctx context.Context) (db *sql.Tx, err error) {
	return r.db.BeginTx(ctx, &sql.TxOptions{})
}

// Commit implements repoContract.
func (r repository) Commit(ctx context.Context, tx *sql.Tx) (err error) {
	return tx.Commit()
}

// CreateAuth implements repoContract.
func (r repository) CreateAuth(ctx context.Context, tx *sql.Tx, auth Auth) (err error) {
	panic("unimplemented")
}

// CreateEmployee implements repoContract.
func (r repository) CreateEmployee(ctx context.Context, tx *sql.Tx, employee Employee) (err error) {
	panic("unimplemented")
}

// FindAuthByEmail implements repoContract.
func (r repository) FindAuthByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `
		SELECT 
			email
		FROM auth
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)
	err = row.Scan(
		&auth.Email,
	)

	if err == sql.ErrNoRows {
		return Auth{}, errEmailNotFound
	} else if err != nil {
		return Auth{}, err
	}

	return
}

// Rollback implements repoContract.
func (r repository) Rollback(ctx context.Context, tx *sql.Tx) (err error) {
	return tx.Rollback()
}

func newRepository(db *sql.DB) repository {
	return repository{db: db}
}
