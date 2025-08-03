package clients

import (
	"fmt"
	"github.com/google/uuid"
)

func (s *session) AddClient(userId uuid.UUID, client *Client) {
	s.mp.Store(userId, client)
}

func (s *session) RemoveClient(userId uuid.UUID) {
	s.mp.Delete(userId)
}

func (s *session) LoadClient(userId uuid.UUID) *Client {
	clientp, ok := s.mp.Load(userId)
	if !ok {
		return nil
	}
	return clientp.(*Client)
}
func (s *session) All() {
	s.mp.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // Продолжаем итерацию
	})
}
