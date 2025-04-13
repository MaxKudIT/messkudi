package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashToPassword(currentvalue string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(currentvalue), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
	}
	return hash
}
func CompareToHash(hash, currentvalue []byte) error {
	if err := bcrypt.CompareHashAndPassword(hash, currentvalue); err != nil {
		return err
	}
	return nil

}
