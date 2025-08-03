package chat

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain/chats"
	"github.com/MaxKudIT/messkudi/internal/domain/clients"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
)

func (csv *chatService) ChatById(ctx context.Context, id uuid.UUID) (dto.ChatDTODetailsServer, error) {
	chat, err := csv.cs.ChatById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			csv.l.Info("chat not found", "error", err)
			return dto.ChatDTODetailsServer{}, err
		}
		csv.l.Info("Error getting chat", "error", err)
		return dto.ChatDTODetailsServer{}, err
	}

	csv.l.Info("Successfully got chat", "id", id)
	return chat, nil

}

func (csv *chatService) ChatByUsersId(ctx context.Context, idone uuid.UUID, idtwo uuid.UUID) (uuid.UUID, error) {
	chatid, err := csv.cs.ChatByUsersId(ctx, idone, idtwo)
	if err != nil {
		csv.l.Error("Error getting chat", "error", err)
		return uuid.Nil, err
	}
	if err == nil && chatid == uuid.Nil {
		csv.l.Info("Chat not found")
		return chatid, nil
	}
	csv.l.Info("Successfully got chat by users", "id", chatid)
	return chatid, nil
}

func (csv *chatService) AllChatsPreview(ctx context.Context, userId uuid.UUID) ([]*dto.ChatPreviewDTOServer, error) {
	previews, err := csv.cs.AllChatsPreview(ctx, userId)
	if err != nil {
		csv.l.Error("Error getting previews", "error", err)
		return nil, err
	}
	if previews == nil {
		return []*dto.ChatPreviewDTOServer{}, nil
	}
	csv.l.Info("Successfully got previews")
	for _, preview := range previews {
		preview.MessageMeta.IsMy = preview.MessageMeta.SenderId == userId
		if id := clients.Session.LoadClient(preview.ParticipantId); id != nil {
			preview.User.Status = true
		}
		preview.MessageMeta.UnReadMessages, err = csv.cms.AllUnreadMessages(ctx, preview.ChatId)
		fmt.Println(preview.MessageMeta.UnReadMessages)
		if err != nil {
			csv.l.Error("Error getting preview", "error", err)
			return nil, err
		}
	}
	return previews, nil
}

func (csv *chatService) CreateChat(ctx context.Context, chat chats.Chat) error {
	if err := csv.cs.SaveChat(ctx, chat); err != nil {

		csv.l.Info("Error saving chat", "error", err)
		return err
	}
	csv.l.Info("Successfully create chat", "id", chat.Id)
	return nil
}

func (csv *chatService) DeleteChat(ctx context.Context, id uuid.UUID) error {
	if err := csv.cs.DeleteChat(ctx, id); err != nil {
		csv.l.Info("Error deleting chat", "error", err)
		return err
	}
	csv.l.Info("Successfully deleted chat", "id", id)
	return nil
}
