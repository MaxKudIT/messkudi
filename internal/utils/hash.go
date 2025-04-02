package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashToPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
	}
	return hash
}
func CompareToHash(hash, password []byte) {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		log.Println(err.Error())
		return
	}

}
