package session

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	User_id   uuid.UUID
	Client_id uuid.UUID
	Device_id string
	Expires   time.Time
	IsActive  bool
}
