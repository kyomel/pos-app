package login

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// GetByEmail implements repoContract.
func (r repository) GetByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `
		SELECT 
			public_id,
			email,
			password,
			is_active,
			role
		FROM auth
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)
	err = row.Scan(
		&auth.PublicId,
		&auth.Email,
		&auth.Password,
		&auth.IsActive,
		&auth.Role,
	)

	if err != nil {
		return
	}

	return
}

func newRepository(db *sql.DB) repository {
	return repository{db: db}
}
