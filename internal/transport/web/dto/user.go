package dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Name        string `validate:"required,min=3,max=50"`
	Password    string `validate:"required,min=6,max=50"`
	PhoneNumber string `validate:"required,regexp=^\\+?[1-9]\\d{1,14}$"`
	Color       string
}

type UserCredentials struct {
	PhoneNumber string
	Password    string
}

type RefreshTokenDTO struct {
	RefreshToken string
}

type UpdateTokenDTO struct {
	Rt RefreshTokenDTO
	Id string
}

func ToDomainWithoutRefresh(id uuid.UUID, createdAt time.Time, updatedAt time.Time, userDTO UserDTO) domain.User {
	return domain.User{
		Id:          id,
		Name:        userDTO.Name,
		Password:    userDTO.Password,
		PhoneNumber: userDTO.PhoneNumber,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

}

func ToDomainUpdate(existsUser domain.User, userDTO UserDTO) domain.UserUpdate {
	result := domain.UserUpdate{Name: existsUser.Name, Password: existsUser.Password, PhoneNumber: existsUser.PhoneNumber, UpdatedAt: time.Now()}
	if userDTO.PhoneNumber != "" {
		result.PhoneNumber = userDTO.PhoneNumber
	}
	if userDTO.Password != "" {
		result.Password = userDTO.Password
	}
	if userDTO.Name != "" {
		result.Name = userDTO.Name
	}
	return result

}
