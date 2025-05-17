package chat_message

import (
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/google/uuid"
	"time"
)

type ChatMessageDTOClient struct {
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //chat
	ChatId             uuid.UUID                //айди чата
}

type ChatMessageDTODetailsServer struct {
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //chat
	ChatId             uuid.UUID                //айди чата
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	ReadAt             *time.Time
	SenderId           uuid.UUID
}
