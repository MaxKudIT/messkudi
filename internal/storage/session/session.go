package session

import (
	"context"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/session"
)

func (ss *sessionStorage) SaveSession(ctx context.Context, session session.Session) error {

	const CREATE_SESSION_QUERY = "INSERT INTO users_sessions (user_id, device_id, expires, isactive) VALUES ($1, $2, $3, $4)"

	if _, err := ss.db.ExecContext(ctx, CREATE_SESSION_QUERY, session.User_id, session.Device_id, session.Expires, session.IsActive); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			ss.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			ss.l.Warn("Query timed out", "error", err)
			return err
		default:
			ss.l.Error("Query failed", "error", err)
			return err
		}
	}
	ss.l.Info("Successfully created session")
	return nil

}
