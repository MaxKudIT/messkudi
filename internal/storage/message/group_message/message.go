package chat_message

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
)

func (cms *chatMessageStorage) MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error) {
	var messagep chat_message_dto.ChatMessageDTODetailsServer
	const GET_CHAT_MESSAGE_QUERY = "SELECT type, content, correspondencetype, sender_id, created_at, updated_at, read_at, chat_id from chat_messages where id = $1"
	if err := cms.db.QueryRowContext(ctx, GET_CHAT_MESSAGE_QUERY, id).Scan(&messagep.Type, &messagep.Content, &messagep.CorrespondenceType, &messagep.SenderId, &messagep.CreatedAt, &messagep.UpdatedAt, &messagep.ReadAt, &messagep.ChatId); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cms.l.Error("user not found", "error", err)
			return chat_message_dto.ChatMessageDTODetailsServer{}, err
		case errors.Is(err, context.Canceled):
			cms.l.Warn("Query cancelled", "error", err)
			return chat_message_dto.ChatMessageDTODetailsServer{}, err
		case errors.Is(err, context.DeadlineExceeded):
			cms.l.Warn("Query timed out", "error", err)
			return chat_message_dto.ChatMessageDTODetailsServer{}, err
		default:
			cms.l.Error("Query failed", "error", err)
			return chat_message_dto.ChatMessageDTODetailsServer{}, err
		}
	}
	cms.l.Info("Successfully got chat message", "id", id)
	return messagep, nil
}

func (cms *chatMessageStorage) SaveMessage(ctx context.Context, message messages.ChatMessage) error {

	const SAVE_CHAT_MESSAGE_QUERY = "INSERT INTO chat_messages (id, type, content, correspondencetype, sender_id, created_at, updated_at, read_at, chat_id ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	if _, err := cms.db.ExecContext(ctx, SAVE_CHAT_MESSAGE_QUERY, message.Id, message.Type, message.Content, message.CorrespondenceType, message.SenderId, message.CreatedAt, message.UpdatedAt, message.ReadAt, message.ChatId); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cms.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			cms.l.Warn("Query timed out", "error", err)
			return err
		default:
			cms.l.Error("Query failed", "error", err)
			return err
		}
	}
	cms.l.Info("Successfully created chat message", "id", message.Id)
	return nil
}

func (cms *chatMessageStorage) DeleteUser(ctx context.Context, id uuid.UUID) error {
	const DELETE_CHAT_MESSAGE_QUERY = "DELETE FROM chat_messages WHERE id = $1"

	if _, err := cms.db.ExecContext(ctx, DELETE_CHAT_MESSAGE_QUERY, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cms.l.Error("user not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			cms.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			cms.l.Warn("Query timed out", "error", err)
			return err
		default:
			cms.l.Error("Query failed", "error", err)
			return err
		}
	}
	cms.l.Info("Successfully deleted chat message", "id", id)
	return nil
}
