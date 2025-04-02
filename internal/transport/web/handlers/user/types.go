package user

import (
	"github.com/MaxKudIT/messkudi/internal/domain/user"
	"github.com/MaxKudIT/messkudi/internal/dto"
	"github.com/google/uuid"
)

type userservice interface {
	CreateUser(userCr dto.UserDTO) (user.User, error)
	GetUser(id uuid.UUID) (user.User, error)
	//UpdateUser(c *gin.Context)
	DeleteUser(id uuid.UUID) error
}

type userhandler struct {
	us userservice
}

func New(us userservice) *userhandler {
	return &userhandler{us}
}
