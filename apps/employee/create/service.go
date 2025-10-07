package create

import (
	"context"
	"database/sql"
)

type repoContract interface {
	FindAuthByEmail(ctx context.Context, email string) (auth Auth, err error)
	CreateAuth(ctx context.Context, tx *sql.Tx, auth Auth) (err error)

	CreateEmployee(ctx context.Context, tx *sql.Tx, employee Employee) (err error)

	// db transaction
	Begin(ctx context.Context) (db *sql.Tx, err error)
	Commit(ctx context.Context, tx *sql.Tx) (err error)
	Rollback(ctx context.Context, tx *sql.Tx) (err error)
}

type service struct {
	repo repoContract
}

func newService(repo repoContract) service {
	return service{repo: repo}
}

func (s service) Create(ctx context.Context, req CreateEmployeeRequest) (err error) {
	// check data auth is exists
	auth, err := s.repo.FindAuthByEmail(ctx, req.Email)
	if err != nil {
		if err != errEmailNotFound {
			return err
		}
	}

	// check auth is exists
	if auth.IsExists() {
		return errEmailAlreadyExist
	}

	// start db transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return err
	}
	defer s.repo.Rollback(ctx, tx)

	return
}
