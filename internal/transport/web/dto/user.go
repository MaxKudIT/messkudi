package dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/google/uuid"
)

type UserDTO struct {
	Name        string `validate:"required,min=3,max=50"`
	LastName    string `validate:"required,min=3,max=50"`
	Password    string `validate:"required,min=6,max=50"`
	PhoneNumber string `validate:"required,regexp=^\\+?[1-9]\\d{1,14}$"`
}

type UserCredentials struct {
	PhoneNumber string
	Password    string
}

type RefreshTokenDTO struct {
	RefreshToken string
}

func ToDomain(id uuid.UUID, createdAt string, updatedAt string, token domain.Token, expiredAt string, userDTO UserDTO) domain.User {
	return domain.User{
		Id:          id,
		Name:        userDTO.Name,
		LastName:    userDTO.LastName,
		Password:    userDTO.Password,
		PhoneNumber: userDTO.PhoneNumber,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Token:       token,
		ExpiredAt:   expiredAt,
	}

}
