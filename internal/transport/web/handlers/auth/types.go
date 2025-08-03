package auth

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/domain/session"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"log/slog"
)

type authService interface {
	UserAuthData(ctx context.Context, usercr dto.UserCredentials) (domain.AuthData, error)
	AccessTokenUpdate(ctx context.Context, data dto.UpdateTokenDTO) (domain.AccessTokenUpdateData, error)
}

type sessionService interface {
	CreateSession(ctx context.Context, session session.Session) error
}

type authHandler struct {
	as authService
	ss sessionService
	l  *slog.Logger
}

func New(as authService, ss sessionService, l *slog.Logger) *authHandler {
	return &authHandler{as: as, ss: ss, l: l}
}
