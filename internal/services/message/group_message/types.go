package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"log/slog"
)

type chatMessageStorage interface {
	MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error)
	SaveMessage(ctx context.Context, message messages.ChatMessage) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type chatMessageService struct {
	cms chatMessageStorage
	l   *slog.Logger
}

func New(cms chatMessageStorage, l *slog.Logger) *chatMessageService {
	return &chatMessageService{cms: cms, l: l}
}
