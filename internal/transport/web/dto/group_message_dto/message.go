package group_message

import (
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/google/uuid"
	"time"
)

type GroupMessageDTOClient struct {
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //group
	ChatId             uuid.UUID                //айди группы
}

type GroupMessageDTODetailsServer struct {
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //group
	ChatId             uuid.UUID                //айди группы
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	ReadAt             *time.Time
}
