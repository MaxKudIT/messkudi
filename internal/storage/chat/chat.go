package chat

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/chats"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
)

func (cs *chatStorage) ChatById(ctx context.Context, id uuid.UUID) (dto.ChatDTODetailsServer, error) {
	var chat dto.ChatDTODetailsServer
	const GET_CHAT_QUERY = "SELECT participant, creator_id, chat_count, createdat, updatedat from chats where id = $1"
	if err := cs.db.QueryRowContext(ctx, GET_CHAT_QUERY, id).Scan(&chat.Participant, &chat.CreatorId, &chat.ChatCount, &chat.CreatedAt, &chat.UpdatedAt); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cs.l.Error("chat not found", "error", err)
			return dto.ChatDTODetailsServer{}, err
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return dto.ChatDTODetailsServer{}, err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return dto.ChatDTODetailsServer{}, err
		default:
			cs.l.Error("Query failed", "error", err)
			return dto.ChatDTODetailsServer{}, err
		}
	}
	cs.l.Info("Successfully got chat", "id", id)
	return chat, nil
}

func (cs *chatStorage) ChatByUsersId(ctx context.Context, idone uuid.UUID, idtwo uuid.UUID) (uuid.UUID, error) {
	var chatid uuid.UUID
	const GetChatQuery = "SELECT id FROM chats WHERE (creator_id = $1 AND participant = $2) OR (creator_id = $2 AND participant = $1)"
	if err := cs.db.QueryRowContext(ctx, GetChatQuery, idone, idtwo).Scan(&chatid); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cs.l.Info("chat not found")
			return uuid.Nil, nil
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return uuid.Nil, err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return uuid.Nil, err
		default:
			cs.l.Error("Query failed", "error", err)
			return uuid.Nil, err
		}
	}
	cs.l.Info("chat exists", "id")
	return chatid, nil
}

func (cs *chatStorage) ChatIsExistsById(ctx context.Context, id uuid.UUID) (bool, error) {
	var isExists bool
	const GetChatIsExist = "SELECT EXISTS (SELECT 1 FROM chats WHERE id = $1);"
	if err := cs.db.QueryRowContext(ctx, GetChatIsExist, id).Scan(&isExists); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return false, err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return false, err
		default:
			cs.l.Error("Query failed", "error", err)
			return false, err
		}
	}
	if isExists {
		cs.l.Info("Chat already exists", "id", id)
	} else {
		cs.l.Info("", "Chat not found by id", id)
	}
	return isExists, nil
}

func (cs *chatStorage) AllChatsPreview(ctx context.Context, userid uuid.UUID) ([]*dto.ChatPreviewDTOServer, error) {
	var check sql.NullTime
	var result []*dto.ChatPreviewDTOServer
	const AllMessagesQuery = `
      WITH last_messages AS (
        SELECT 
            cm.sender_id,
            cm.chat_id,
            cm.content,
            cm.created_at,
            cm.read_at,
            ROW_NUMBER() OVER (PARTITION BY cm.chat_id ORDER BY cm.created_at DESC) as rn
        FROM chat_messages cm
        JOIN chats c ON c.id = cm.chat_id
        WHERE c.participant = $1 OR c.creator_id = $1
    )
    SELECT 
        c.id,
        lm.sender_id,
        lm.content,
        lm.created_at,
        lm.read_at,
        u.name,
        u.color,
    	u.id
    FROM chats c
    JOIN last_messages lm ON c.id = lm.chat_id AND lm.rn = 1
    JOIN users u ON u.id = CASE 
        WHEN c.participant = $1 THEN c.creator_id
        ELSE c.participant
    END
    WHERE c.participant = $1 OR c.creator_id = $1
    ORDER BY lm.created_at DESC`

	rows, err := cs.db.QueryContext(ctx, AllMessagesQuery, userid)
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return nil, err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return nil, err
		default:
			cs.l.Error("Query failed", "error", err)
			return nil, err
		}
	}

	for rows.Next() {
		var currentObject dto.ChatPreviewDTOServer
		if err := rows.Scan(&currentObject.ChatId, &currentObject.MessageMeta.SenderId, &currentObject.MessageMeta.Content, &currentObject.MessageMeta.CreatedAt, &check, &currentObject.User.Name, &currentObject.User.Color, &currentObject.ParticipantId); err != nil {
			cs.l.Error("Scan failed", "error", err)
			return nil, err
		}
		currentObject.MessageMeta.IsRead = check.Valid
		result = append(result, &currentObject)
	}
	cs.l.Info("Successfully getting chat previews")
	return result, nil
}

func (cs *chatStorage) SaveChat(ctx context.Context, chat chats.Chat) error {

	const SAVE_CHAT_QUERY = "INSERT INTO chats (id, participant, creator_id, createdat, updatedat, user_id_one, user_id_two) VALUES ($1, $2::uuid, $3::uuid, $4, $5, LEAST($2::uuid, $3::uuid), GREATEST($2::uuid, $3::uuid))"

	if _, err := cs.db.ExecContext(ctx, SAVE_CHAT_QUERY, chat.Id, chat.Participant, chat.CreatorId, chat.CreatedAt, chat.UpdatedAt); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return err
		default:
			cs.l.Error("Query failed", "error", err)
			return err
		}
	}
	cs.l.Info("Successfully saving chat", "id", chat.Id)
	return nil
}

func (cs *chatStorage) DeleteChat(ctx context.Context, id uuid.UUID) error {
	const DELETE_CHAT_QUERY = "DELETE FROM chats WHERE id = $1"

	if _, err := cs.db.ExecContext(ctx, DELETE_CHAT_QUERY, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cs.l.Error("chat not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return err
		default:
			cs.l.Error("Query failed", "error", err)
			return err
		}
	}
	cs.l.Info("Successfully deleted chat", "id", id)
	return nil
}
