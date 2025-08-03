package chat_message

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"time"
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

func (cms *chatMessageStorage) AllUnreadMessages(ctx context.Context, chatId uuid.UUID) ([]uuid.UUID, error) {
	var unread []uuid.UUID = make([]uuid.UUID, 0)

	const AllUnReadMessagesQuery = "SELECT sender_id from chat_messages where chat_id = $1 AND read_at IS NULL"

	rows, err := cms.db.QueryContext(ctx, AllUnReadMessagesQuery, chatId)
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cms.l.Warn("Query cancelled", "error", err)
			return nil, err
		case errors.Is(err, context.DeadlineExceeded):
			cms.l.Warn("Query timed out", "error", err)
			return nil, err
		default:
			cms.l.Error("Query failed", "error", err)
			return nil, err
		}
	}

	for rows.Next() {
		var currentObject uuid.UUID
		if err := rows.Scan(&currentObject); err != nil {
			cms.l.Error("Scan failed", "error", err)
			return nil, err
		}
		unread = append(unread, currentObject)

	}
	cms.l.Info("Successfully getting messages")
	return unread, nil
}

func (cms *chatMessageStorage) AllMessages(ctx context.Context, chatid uuid.UUID) ([]chat_message_dto.ChatMessageDTODetailsServer, error) {
	var messages []chat_message_dto.ChatMessageDTODetailsServer
	var (
		rda sql.NullTime
	)

	const AllMessagesQuery = "SELECT id, type, content, correspondencetype, sender_id, created_at, updated_at, read_at from chat_messages where chat_id = $1 order by created_at"

	rows, err := cms.db.QueryContext(ctx, AllMessagesQuery, chatid)
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cms.l.Warn("Query cancelled", "error", err)
			return nil, err
		case errors.Is(err, context.DeadlineExceeded):
			cms.l.Warn("Query timed out", "error", err)
			return nil, err
		default:
			cms.l.Error("Query failed", "error", err)
			return nil, err
		}
	}

	for rows.Next() {
		var currentObject chat_message_dto.ChatMessageDTODetailsServer
		if err := rows.Scan(&currentObject.Id, &currentObject.Type, &currentObject.Content, &currentObject.CorrespondenceType, &currentObject.SenderId, &currentObject.CreatedAt, &currentObject.UpdatedAt, &rda); err != nil {
			cms.l.Error("Scan failed", "error", err)
			return nil, err
		}

		if !rda.Valid {
			currentObject.ReadAt = nil
		} else {
			currentObject.ReadAt = &rda.Time
		}
		currentObject.ChatId = chatid
		messages = append(messages, currentObject)
	}
	cms.l.Info("Successfully getting messages")
	return messages, nil
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
	cms.l.Info("Successfully saving chat message", "id", message.Id)
	return nil
}

func (cms *chatMessageStorage) UpdateMessage(ctx context.Context, message chat_message_dto.ChatMessageDTOClientParsing) error {
	const UpdateChatMessageQuery = "UPDATE chat_messages SET content = $1 WHERE chat_id = $2"

	if _, err := cms.db.ExecContext(ctx, UpdateChatMessageQuery, message.Content, message.ChatId); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cms.l.Error("message not found", "error", err)
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
	cms.l.Info("Successfully updated chat message")
	return nil
}

func (cms *chatMessageStorage) UpdateReadAtMessage(ctx context.Context, time time.Time) error {
	const UpdateChatMessageQuery = "UPDATE chat_messages SET read_at = $1 WHERE chat_id = $2"

	if _, err := cms.db.ExecContext(ctx, UpdateChatMessageQuery, time); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cms.l.Error("message not found", "error", err)
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
	cms.l.Info("Successfully updated readat chat message")
	return nil
}

func (cms *chatMessageStorage) DeleteMessage(ctx context.Context, id uuid.UUID) error {
	const DELETE_CHAT_MESSAGE_QUERY = "DELETE FROM chat_messages WHERE id = $1"

	if _, err := cms.db.ExecContext(ctx, DELETE_CHAT_MESSAGE_QUERY, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cms.l.Error("message not found", "error", err)
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
