package chats

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	Id          uuid.UUID
	CreatorId   uuid.UUID
	Participant uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
