package utils

import (
	"errors"
	"github.com/nyaruka/phonenumbers"
)

func ValidatePhone(phone, countryCode string) (string, error) {
	num, err := phonenumbers.Parse(phone, countryCode)
	if err != nil {
		return "", err
	}
	if !phonenumbers.IsValidNumber(num) {
		return "", errors.New("invalid phone number")
	}
	return phonenumbers.Format(num, phonenumbers.E164), nil
}
