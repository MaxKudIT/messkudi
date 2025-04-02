package utils

import "github.com/google/uuid"

func GenerationUUID() uuid.UUID {
	return uuid.New()
}
