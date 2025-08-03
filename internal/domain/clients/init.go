package clients

import "sync"

var (
	instance *session
	once     sync.Once
)

func GetSession() *session {
	once.Do(func() {
		instance = NewSession(&sync.Map{})
	})
	return instance
}

var Session *session = &session{mp: sync.Map{}}
