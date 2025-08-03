package chat

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/chats"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type chatStorage interface {
	ChatById(ctx context.Context, id uuid.UUID) (dto.ChatDTODetailsServer, error)
	SaveChat(ctx context.Context, chat chats.Chat) error
	AllChatsPreview(ctx context.Context, userid uuid.UUID) ([]*dto.ChatPreviewDTOServer, error)
	ChatByUsersId(ctx context.Context, idone uuid.UUID, idtwo uuid.UUID) (uuid.UUID, error)
	DeleteChat(ctx context.Context, id uuid.UUID) error
}

type chatMessageStorage interface {
	MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error)
	AllMessages(ctx context.Context, chatid uuid.UUID) ([]chat_message_dto.ChatMessageDTODetailsServer, error)
	SaveMessage(ctx context.Context, message messages.ChatMessage) error
	UpdateMessage(ctx context.Context, message chat_message_dto.ChatMessageDTOClientParsing) error
	UpdateReadAtMessage(ctx context.Context, time time.Time) error
	AllUnreadMessages(ctx context.Context, chatId uuid.UUID) ([]uuid.UUID, error)
	DeleteMessage(ctx context.Context, id uuid.UUID) error
}
type chatService struct {
	cs  chatStorage
	cms chatMessageStorage
	l   *slog.Logger
}

func New(cs chatStorage, cms chatMessageStorage, l *slog.Logger) *chatService {
	return &chatService{cs: cs, cms: cms, l: l}
}
