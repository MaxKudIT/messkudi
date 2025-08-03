package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/group_message_dto"
	"github.com/google/uuid"
	"log/slog"
)

type groupMessageService interface {
	MessageById(ctx context.Context, id uuid.UUID) (group_message_dto.GroupMessageDTODetailsServer, error)
	CreateMessage(ctx context.Context, message messages.GroupMessage) error
	DeleteMessage(ctx context.Context, id uuid.UUID) error
}

type groupMessageHandler struct {
	gmsv groupMessageService
	l    *slog.Logger
}

func New(gmsv groupMessageService, l *slog.Logger) *groupMessageHandler {
	return &groupMessageHandler{gmsv: gmsv, l: l}
}
