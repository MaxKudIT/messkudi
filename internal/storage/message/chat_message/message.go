package message

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
)

func (ms *messageStorage) MessageById(ctx context.Context, id uuid.UUID) (dto.MessageDTODetailsServer, error) {
	var messagep dto.MessageDTODetailsServer
	const GET_MESSAGE_QUERY = "SELECT type, COALESCE(chat_id, group_id) AS correspondence_id, correspondencetype, content, created_at, updated_at, read_at FROM Messages WHERE id = $1"
	if err := ms.db.QueryRowContext(ctx, GET_MESSAGE_QUERY, id).Scan(&messagep.Type, &messagep.CorrespondenceId, &messagep.CorrespondenceType, &messagep.Content, &messagep.CreatedAt, &messagep.UpdatedAt, &messagep.ReadAt); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			ms.l.Error("user not found", "error", err)
			return dto.MessageDTODetailsServer{}, err
		case errors.Is(err, context.Canceled):
			ms.l.Warn("Query cancelled", "error", err)
			return dto.MessageDTODetailsServer{}, err
		case errors.Is(err, context.DeadlineExceeded):
			ms.l.Warn("Query timed out", "error", err)
			return dto.MessageDTODetailsServer{}, err
		default:
			ms.l.Error("Query failed", "error", err)
			return dto.MessageDTODetailsServer{}, err
		}
	}
	ms.l.Info("Successfully got message", "id", id)
	return messagep, nil
}

func (ms *messageStorage) SaveMessage(ctx context.Context, message messages.Message) error {

	const SAVE_MESSAGE_QUERY = "INSERT INTO messages (id, content, createdat, updatedat, readat, sender_id, group_id, chat_id, type) VALUES ($1, $2, $3, $4, $5, $6, case when &7 is not null then &7 end, case when &8 is not null then &8 end, $9)"

	if _, err := ms.db.ExecContext(ctx, SAVE_MESSAGE_QUERY, message.Id, message.Content, message.CreatedAt, message.UpdatedAt, message.ReadAt, message.SenderId, userp.Token.RefreshToken, userp.ExpiredAt); err != nil {
		switch { //gggggggg
		case errors.Is(err, context.Canceled):
			us.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			us.l.Warn("Query timed out", "error", err)
			return err
		default:
			us.l.Error("Query failed", "error", err)
			return err
		}
	}
	us.l.Info("Successfully created user", "id", userp.Id)
	return nil
}

func (us *userStorage) DeleteUser(ctx context.Context, id uuid.UUID) error {
	const DELETE_USER_QUERY = "DELETE FROM users WHERE id = $1"

	if _, err := us.db.ExecContext(ctx, DELETE_USER_QUERY, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			us.l.Error("user not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			us.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			us.l.Warn("Query timed out", "error", err)
			return err
		default:
			us.l.Error("Query failed", "error", err)
			return err
		}
	}
	us.l.Info("Successfully deleted user", "id", id)
	return nil
}
