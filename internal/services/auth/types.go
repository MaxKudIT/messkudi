package auth

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type authStorage interface {
	UserAuthData(ctx context.Context, usercr dto.UserCredentials) (domain.AuthData, error)
	UserUpdateRefreshToken(ctx context.Context, newRt string, expired time.Time, pn string) error
	AccessTokenUpdate(ctx context.Context, id uuid.UUID) (domain.AccessTokenUpdateData, error)
}

type authService struct {
	ast authStorage
	l   *slog.Logger
}

func New(ast authStorage, l *slog.Logger) *authService {
	return &authService{ast: ast, l: l}
}
