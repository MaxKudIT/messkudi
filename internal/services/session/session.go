package session

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/session"
)

func (u *sessionService) CreateSession(ctx context.Context, session session.Session) error { //DOMAIN

	if err := u.ss.SaveSession(ctx, session); err != nil {
		u.l.Error("Error saving session", "error", err)
		return err
	}
	u.l.Info("Successfully saving session")
	return nil
}
