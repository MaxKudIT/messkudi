package clients

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	Conn     *websocket.Conn
	ClientId uuid.UUID
}

type session struct {
	mp sync.Map
}

func NewSession(mp *sync.Map) *session {
	return &session{mp: *mp}
}

//ограничение, jsonb
