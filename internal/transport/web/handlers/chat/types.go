package chat

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/domain/chats"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"log/slog"
)

type chatService interface {
	ChatById(ctx context.Context, id uuid.UUID) (dto.ChatDTODetailsServer, error)
	ChatByUsersId(ctx context.Context, idone uuid.UUID, idtwo uuid.UUID) (uuid.UUID, error)
	AllChatsPreview(ctx context.Context, userId uuid.UUID) ([]*dto.ChatPreviewDTOServer, error)
	CreateChat(ctx context.Context, chat chats.Chat) error
	DeleteChat(ctx context.Context, id uuid.UUID) error
}

type userService interface {
	CreateUser(ctx context.Context, userp domain.User) (domain.User, error)
	UserById(ctx context.Context, id uuid.UUID) (dto.UserDTO, error)
	UserByPhoneNumber(ctx context.Context, phoneNumber string) (dto.UserDTO, error)
	UserIdByPhoneNumber(ctx context.Context, phoneNumber string) (uuid.UUID, error)
	UserIsExistsByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error)
	UserDataForChatHeader(ctx context.Context, id uuid.UUID) (dto.ChatHeader, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type chatMessageService interface {
	MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error)
	AllMessages(ctx context.Context, chatid uuid.UUID) ([]chat_message_dto.ChatMessageDTODetailsServer, error)
	CreateMessage(ctx context.Context, message messages.ChatMessage) error
	DeleteMessage(ctx context.Context, id uuid.UUID) error
}
type chatHandler struct {
	csv chatService
	cms chatMessageService
	us  userService
	l   *slog.Logger
}

func New(csv chatService, cms chatMessageService, us userService, l *slog.Logger) *chatHandler {
	return &chatHandler{csv: csv, cms: cms, us: us, l: l}
}
