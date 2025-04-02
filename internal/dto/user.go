package dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain/auth"
	"github.com/MaxKudIT/messkudi/internal/domain/user"
	"github.com/google/uuid"
)

type UserDTO struct {
	Name        string `validate:"required,min=3,max=50"`
	LastName    string `validate:"required,min=3,max=50"`
	Password    string `validate:"required,min=6,max=50"`
	PhoneNumber string `validate:"required,regexp=^\\+?[1-9]\\d{1,14}$"`
}

func ToDomain(id uuid.UUID, createdAt string, updatedAt string, token auth.Token, userDTO UserDTO) user.User {
	return user.User{
		Id:          id,
		Name:        userDTO.Name,
		LastName:    userDTO.LastName,
		Password:    userDTO.Password,
		PhoneNumber: userDTO.PhoneNumber,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Token:       token,
	}

}
