package chat_message

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/group_message_dto"
	"github.com/google/uuid"
)

func (gmsv *groupMessageService) MessageById(ctx context.Context, id uuid.UUID) (group_message_dto.GroupMessageDTODetailsServer, error) {
	message, err := gmsv.gms.MessageById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			gmsv.l.Error("Message not found", "error", err)
			return group_message_dto.GroupMessageDTODetailsServer{}, err
		}
		gmsv.l.Error("Error getting group message", "error", err)
		return group_message_dto.GroupMessageDTODetailsServer{}, err
	}
	gmsv.l.Info("Successfully getting group message", "id", id)
	return message, nil
}

func (gmsv *groupMessageService) CreateMessage(ctx context.Context, message messages.GroupMessage) error {
	if err := gmsv.gms.SaveMessage(ctx, message); err != nil {
		gmsv.l.Error("Error creating group message", "error", err)
		return err
	}
	gmsv.l.Info("Successfully created group message", "id", message.Id)
	return nil
}

func (gmsv *groupMessageService) DeleteMessage(ctx context.Context, id uuid.UUID) error {
	if err := gmsv.gms.DeleteUser(ctx, id); err != nil {
		gmsv.l.Error("Error deleting group message", "error", err)
		return err
	}
	gmsv.l.Info("Successfully deleted group message", "id", id)
	return nil
}
