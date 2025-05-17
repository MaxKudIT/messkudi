package websocket

import "github.com/google/uuid"

type Connection struct {
	userid  uuid.UUID
	message string
}
type Disconnection struct {
	userid  uuid.UUID
	message string
}
