package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID
	Name        string    `validate:"required,min=3,max=50"`
	Password    string    `validate:"required,min=6,max=50"`
	PhoneNumber string    `validate:"required,regexp=^\\+?[1-9]\\d{1,14}$"`
	CreatedAt   time.Time //string `json:"created_at"` //в time.Time
	UpdatedAt   time.Time //string `json:"updated_at"`
	Token       Token
	ExpiredAt   time.Time
	Color       string
}

type UserUpdate struct {
	Name        string    `validate:"required,min=3,max=50"`
	Password    string    `validate:"required,min=6,max=50"`
	PhoneNumber string    `validate:"required,regexp=^\\+?[1-9]\\d{1,14}$"`
	UpdatedAt   time.Time //string `json:"updated_at"`
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
	Id          uuid.UUID
	Name        string
	Password    string
	PhoneNumber string
	Token       Token
}
