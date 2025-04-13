package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID
	Name        string `validate:"required,min=3,max=50"`
	LastName    string `validate:"required,min=3,max=50"`
	Password    string `validate:"required,min=6,max=50"`
	PhoneNumber string `validate:"required,regexp=^\\+?[1-9]\\d{1,14}$"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Token       Token
	ExpiredAt   string
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenUpdateData struct {
	Id           uuid.UUID
	AccessToken  string
	RefreshToken string
}

type AuthData struct {
	Id       uuid.UUID
	Name     string
	Password string
	Token    Token
}
