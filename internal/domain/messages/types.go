package messages

import (
	"github.com/google/uuid"
	"time"
)

type CorrespondenceT string
type MessageT string

const (
	Chat    CorrespondenceT = "chat"
	Group   CorrespondenceT = "group"
	Channel CorrespondenceT = "channel"
)
const (
	Text  MessageT = "text"
	Image MessageT = "image"
	File  MessageT = "file"
	Link  MessageT = "link"
	Audio MessageT = "audio"
	Video MessageT = "video"
)

type ChatMessage struct {
	Id                 uuid.UUID
	Type               MessageT
	Content            string
	CorrespondenceType CorrespondenceT //chat или group
	ChatId             uuid.UUID       //айди чата или группы
	SenderId           uuid.UUID
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ReadAt             *time.Time
}

type GroupMessage struct {
	Id                 uuid.UUID
	Type               MessageT
	Content            string
	CorrespondenceType CorrespondenceT //chat или group
	GroupId            uuid.UUID       //айди чата или группы
	SenderId           uuid.UUID
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ReadAt             *time.Time
}
