package create

import (
	"context"
	"database/sql"
	"log/slog"
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

func (s service) create(ctx context.Context, req CreateEmployeeRequest) (err error) {
	// check data auth is exists
	auth, err := s.repo.FindAuthByEmail(ctx, req.Email)
	if err != nil {
		if err != errEmailNotFound {
			slog.ErrorContext(ctx, "[create] error when FindAuthByEmail", slog.String("error", err.Error()))
			return err
		}
	}

	// check auth is exists
	if auth.IsExists() {
		slog.ErrorContext(ctx, "[create] email is exists")
		return errEmailAlreadyExist
	}

	// start db transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "[create] error when start db transaction", slog.Any("error", err.Error()))
		return err
	}
	defer s.repo.Rollback(ctx, tx)

	authModel := req.ToAuthModel()
	if err := s.repo.CreateAuth(ctx, tx, authModel); err != nil {
		slog.ErrorContext(ctx, "[create] error when try to CreateAuth", slog.Any("error", err.Error()))
		return err
	}

	empModel := req.ToEmployeeModel(auth.PublicId)
	if err := s.repo.CreateEmployee(ctx, tx, empModel); err != nil {
		slog.ErrorContext(ctx, "[create] error when try to CreateEmployee", slog.Any("error", err.Error()))
		return err
	}

	err = s.repo.Commit(ctx, tx)
	if err != nil {
		slog.ErrorContext(ctx, "[create] error when try to commit transaction", slog.Any("error", err.Error()))
		return err
	}

	return
}
