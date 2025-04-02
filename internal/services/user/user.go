package user

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain/auth"
	"github.com/MaxKudIT/messkudi/internal/domain/user"
	"github.com/MaxKudIT/messkudi/internal/dto"
	"github.com/MaxKudIT/messkudi/internal/utils"
	"github.com/google/uuid"
	"time"
)

func (u *userService) CreateUser(userCr dto.UserDTO) (user.User, error) {
	id := utils.GenerationUUID()
	date := time.Now().Format("02-01-06")
	hash := utils.HashToPassword(userCr.Password)
	hashhex := hex.EncodeToString(hash)
	refreshToken := utils.CreateRefreshToken()
	refreshTokenhex := hex.EncodeToString(refreshToken)
	userCrNew := dto.UserDTO{Name: userCr.Name, LastName: userCr.LastName, Password: hashhex, PhoneNumber: userCr.PhoneNumber}
	userp := dto.ToDomain(id, date, date, auth.Token{RefreshToken: refreshTokenhex}, userCrNew)
	if err := u.us.SaveUser(userp); err != nil {
		return user.User{}, fmt.Errorf("failed to save user")
	}
	return userp, nil
}

func (u *userService) GetUser(id uuid.UUID) (user.User, error) {
	userp, err := u.us.GetUserById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}, fmt.Errorf("user not found")
		} else {
			return user.User{}, fmt.Errorf("failed to fetch user")
		}
	}
	return userp, nil
}

//func (userServ UserServices) UpdateUserService(userUpd user.UserUpdate) {
//	userServ.userRepo.UpdateUser(userUpd)
//}

func (u *userService) DeleteUser(id uuid.UUID) error {
	if err := u.us.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
