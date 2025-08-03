package chat_message

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
)

func (cmsv *chatMessageService) MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error) {
	message, err := cmsv.cms.MessageById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cmsv.l.Error("Message not found", "error", err)
			return chat_message_dto.ChatMessageDTODetailsServer{}, err
		}
		cmsv.l.Error("Error getting chat message", "error", err)
		return chat_message_dto.ChatMessageDTODetailsServer{}, err
	}
	cmsv.l.Info("Successfully getting chat message", "id", id)
	return chat_message_dto.ChatMessageDTODetailsServer(message), nil
}

func (cmsv *chatMessageService) AllMessages(ctx context.Context, chatid uuid.UUID) ([]chat_message_dto.ChatMessageDTODetailsServer, error) {
	messages, err := cmsv.cms.AllMessages(ctx, chatid)
	if err != nil {
		cmsv.l.Error("Error getting all messages", "error", err)
		return nil, err
	}
	cmsv.l.Info("Successfully getting all messages", "id", chatid)
	return messages, nil
}

func (cmsv *chatMessageService) CreateMessage(ctx context.Context, message messages.ChatMessage) error {
	if err := cmsv.cms.SaveMessage(ctx, message); err != nil {
		cmsv.l.Error("Error creating chat message", "error", err)
		return err
	}
	cmsv.l.Info("Successfully created chat message", "id", message.Id)
	return nil
}

func (cmsv *chatMessageService) DeleteMessage(ctx context.Context, id uuid.UUID) error {
	if err := cmsv.cms.DeleteMessage(ctx, id); err != nil {
		cmsv.l.Error("Error deleting chat message", "error", err)
		return err
	}
	cmsv.l.Info("Successfully deleted chat message", "id", id)
	return nil
}
