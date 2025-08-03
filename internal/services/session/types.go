package session

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/session"
	"log/slog"
)

type sessionStorage interface {
	SaveSession(ctx context.Context, session session.Session) error
}

type sessionService struct {
	ss sessionStorage
	l  *slog.Logger
}

func New(ss sessionStorage, l *slog.Logger) *sessionService {
	return &sessionService{ss: ss, l: l}
}
