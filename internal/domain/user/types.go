package user

import (
	"github.com/MaxKudIT/messkudi/internal/domain/auth"
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
	Token       auth.Token
}
