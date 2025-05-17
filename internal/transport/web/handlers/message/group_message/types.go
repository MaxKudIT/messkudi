package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"log/slog"
)

type chatMessageService interface {
	MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error)
	CreateMessage(ctx context.Context, message messages.ChatMessage) error
	DeleteMessage(ctx context.Context, id uuid.UUID) error
}

type chatMessageHandler struct {
	cmsv chatMessageService
	l    *slog.Logger
}

func New(cmsv chatMessageService, l *slog.Logger) *chatMessageHandler {
	return &chatMessageHandler{cmsv: cmsv, l: l}
}
