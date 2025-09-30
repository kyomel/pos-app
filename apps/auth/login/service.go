package login

import (
	"context"
	"log/slog"
	"time"

	"github.com/kyomel/pos-app/internal/config"
	token "github.com/kyomel/pos-app/internal/utils/jwt"
)

type repoContract interface {
	GetByEmail(ctx context.Context, email string) (auth Auth, err error)
}

type service struct {
	repo repoContract
}

func newService(repo repoContract) service {
	return service{repo: repo}
}

func (s *service) login(ctx context.Context, req LoginRequest) (tok string, role string, err error) {
	auth, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		slog.ErrorContext(ctx, "[login] error when try to GetByEmail", slog.Any("error", err.Error()))
		return
	}

	if !auth.IsActive {
		slog.ErrorContext(ctx, "[login] account is not active", slog.Any("email", auth.Email))
		return "", "", errAccountIsNotActive
	}

	if err = auth.ValidatePassword(req.Password); err != nil {
		slog.ErrorContext(ctx, "[login] error when try to validate password", slog.Any("error", err.Error()))
		return "", "", err
	}

	// kemudian, kita buat token JWT
	cfg := config.GetConfig()
	claims := token.Claims{
		Id:         auth.PublicId,
		Role:       auth.Role,
		ExpireTime: time.Duration(cfg.App.ExpireTime * int(time.Second)),
	}

	// setelah itu, kita generate token JWT
	token, err := auth.GenerateToken(claims, cfg.App.SecretKey)
	if err != nil {
		slog.ErrorContext(ctx, "[login] error when create token", slog.Any("error", err.Error()))
		return "", "", err
	}

	return token, auth.Role, nil
}
