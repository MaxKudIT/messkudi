package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"math/rand"
)

func GenerationUUID() uuid.UUID {
	return uuid.New()
}

func GenerationDeviceHash(Ip string, useragent string) string {

	combinedString := Ip + ":" + useragent

	hash := sha256.Sum256([]byte(combinedString))

	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func GetRandomColor() string {
	colors := []string{
		"orange",
		"purple",
		"green",
		"red",
		"blue",
		"coral",
		"lightsalmon",
	}

	randomIndex := rand.Intn(len(colors))

	return colors[randomIndex]
}
