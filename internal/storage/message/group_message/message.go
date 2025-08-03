package chat_message

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/group_message_dto"
	"github.com/google/uuid"
)

func (gms *groupMessageStorage) MessageById(ctx context.Context, id uuid.UUID) (group_message_dto.GroupMessageDTODetailsServer, error) {
	var messagep group_message_dto.GroupMessageDTODetailsServer
	const GET_GROUP_MESSAGE_QUERY = "SELECT type, content, correspondencetype, sender_id, created_at, updated_at, read_at, group_id from group_messages where id = $1"
	if err := gms.db.QueryRowContext(ctx, GET_GROUP_MESSAGE_QUERY, id).Scan(&messagep.Type, &messagep.Content, &messagep.CorrespondenceType, &messagep.SenderId, &messagep.CreatedAt, &messagep.UpdatedAt, &messagep.ReadAt, &messagep.GroupId); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			gms.l.Error("user not found", "error", err)
			return group_message_dto.GroupMessageDTODetailsServer{}, err
		case errors.Is(err, context.Canceled):
			gms.l.Warn("Query cancelled", "error", err)
			return group_message_dto.GroupMessageDTODetailsServer{}, err
		case errors.Is(err, context.DeadlineExceeded):
			gms.l.Warn("Query timed out", "error", err)
			return group_message_dto.GroupMessageDTODetailsServer{}, err
		default:
			gms.l.Error("Query failed", "error", err)
			return group_message_dto.GroupMessageDTODetailsServer{}, err
		}
	}
	gms.l.Info("Successfully got group message", "id", id)
	return messagep, nil
}

func (gms *groupMessageStorage) SaveMessage(ctx context.Context, message messages.GroupMessage) error {

	const SAVE_GROUP_MESSAGE_QUERY = "INSERT INTO group_messages (id, type, content, correspondencetype, sender_id, created_at, updated_at, read_at, group_id ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	if _, err := gms.db.ExecContext(ctx, SAVE_GROUP_MESSAGE_QUERY, message.Id, message.Type, message.Content, message.CorrespondenceType, message.SenderId, message.CreatedAt, message.UpdatedAt, message.ReadAt, message.GroupId); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			gms.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			gms.l.Warn("Query timed out", "error", err)
			return err
		default:
			gms.l.Error("Query failed", "error", err)
			return err
		}
	}
	gms.l.Info("Successfully created group message", "id", message.Id)
	return nil
}

func (gms *groupMessageStorage) DeleteUser(ctx context.Context, id uuid.UUID) error {
	const DELETE_GROUP_MESSAGE_QUERY = "DELETE FROM group_messages WHERE id = $1"

	if _, err := gms.db.ExecContext(ctx, DELETE_GROUP_MESSAGE_QUERY, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			gms.l.Error("user not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			gms.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			gms.l.Warn("Query timed out", "error", err)
			return err
		default:
			gms.l.Error("Query failed", "error", err)
			return err
		}
	}
	gms.l.Info("Successfully deleted group message", "id", id)
	return nil
}
