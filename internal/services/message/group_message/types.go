package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/group_message_dto"
	"github.com/google/uuid"
	"log/slog"
)

type groupMessageStorage interface {
	MessageById(ctx context.Context, id uuid.UUID) (group_message_dto.GroupMessageDTODetailsServer, error)
	SaveMessage(ctx context.Context, message messages.GroupMessage) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type groupMessageService struct {
	gms groupMessageStorage
	l   *slog.Logger
}

func New(gms groupMessageStorage, l *slog.Logger) *groupMessageService {
	return &groupMessageService{gms: gms, l: l}
}
