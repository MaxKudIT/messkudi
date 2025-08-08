package websocket

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

type ChatMessageStorage interface {
	MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error)
	AllMessages(ctx context.Context, chatid uuid.UUID) ([]chat_message_dto.ChatMessageDTODetailsServer, error)
	SaveMessage(ctx context.Context, message messages.ChatMessage) error
	UpdateMessage(ctx context.Context, message chat_message_dto.ChatMessageDTOClientParsing) error
	UpdateReadAtMessage(ctx context.Context, time time.Time, messageId uuid.UUID) error
	AllUnreadMessages(ctx context.Context, chatId uuid.UUID) ([]uuid.UUID, error)
	DeleteMessage(ctx context.Context, id uuid.UUID) error
}

type ChatStorage interface {
	ChatById(ctx context.Context, id uuid.UUID) (dto.ChatDTODetailsServer, error)
	ChatByUsersId(ctx context.Context, idone uuid.UUID, idtwo uuid.UUID) (uuid.UUID, error)
	ChatIsExistsById(ctx context.Context, id uuid.UUID) (bool, error)
	AllChatsPreview(ctx context.Context, userid uuid.UUID) ([]*dto.ChatPreviewDTOServer, error)
	SaveChat(ctx context.Context, chat chats.Chat) error
	DeleteChat(ctx context.Context, id uuid.UUID) error
}

type rabbitmq interface {
	ConsumeMessages(queueName string) ([]chat_message_dto.ChatMessageDTOClient, error)
	ProduceMessage(data chat_message_dto.ChatMessageDTOClient) error
	Setup(queueName string) error
}

type websocketService struct {
	mailing chan chat_message_dto.ChatMessageDTOClientParsing
	sid     uuid.UUID
	cms     ChatMessageStorage
	cs      ChatStorage
	rabbit  rabbitmq
	l       *slog.Logger
}

func New(sid uuid.UUID, mailing chan chat_message_dto.ChatMessageDTOClientParsing, cms ChatMessageStorage, cs ChatStorage, rabbit rabbitmq, logger *slog.Logger) *websocketService {
	return &websocketService{sid: sid, mailing: mailing, rabbit: rabbit, cms: cms, cs: cs, l: logger}
}
