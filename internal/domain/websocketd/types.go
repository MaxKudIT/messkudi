package websocketd

import (
	"github.com/google/uuid"
)

type Connection struct {
	Userid  uuid.UUID
	Message string
}
type Disconnection struct {
	Userid  uuid.UUID
	Message string
}
