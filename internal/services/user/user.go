package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/domain/clients"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/utils"
	"time"

	"github.com/google/uuid"
)

func (u *userService) CreateUser(ctx context.Context, userp domain.User) (domain.User, error) { //DOMAIN
	userp.Color = utils.GetRandomColor()
	if err := u.us.SaveUser(ctx, userp); err != nil {
		u.l.Error("Error saving user", "error", err)
		return domain.User{}, err
	}
	u.l.Info("Successfully saving user")
	return userp, nil
}

func (u *userService) UserById(ctx context.Context, id uuid.UUID) (dto.UserDTO, error) {

	userp, err := u.us.UserById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.l.Error("User not found", "id", id)
			return dto.UserDTO{}, err
		} else {
			u.l.Error("Error getting user", "error", err)
			return dto.UserDTO{}, err
		}
	}
	u.l.Info("Successfully fetched user", "id", id)
	return userp, nil
}

//func (userServ UserServices) UpdateUser(userUpd user.UserUpdate) {
//	userServ.userRepo.UpdateUser(userUpd)
//}

func (u *userService) UserByPhoneNumber(ctx context.Context, phoneNumber string) (dto.UserDTO, error) {
	userp, err := u.us.UserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.l.Error("User not found", "phonenumber", phoneNumber)
			return dto.UserDTO{}, err
		} else {
			u.l.Error("Error getting user", "error", err)
			return dto.UserDTO{}, err
		}
	}
	u.l.Info("Successfully fetched user", "phonenumber", phoneNumber)
	return userp, nil
}

func (u *userService) UserIdByPhoneNumber(ctx context.Context, phoneNumber string) (uuid.UUID, error) {
	id, err := u.us.UserIdByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.l.Error("User not found", "phonenumber", phoneNumber)
			return uuid.Nil, err
		} else {
			u.l.Error("Error getting user id", "error", err)
			return uuid.Nil, err
		}
	}
	u.l.Info("Successfully fetched user id", "phonenumber", phoneNumber)
	return id, nil
}

func (u *userService) UserIsExistsByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error) {
	isExists, err := u.us.UserIsExistsByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		u.l.Error("Error getting result isExistsByPN", "phonenumber", phoneNumber)
		return false, err
	}
	u.l.Info("Successfully getting result isExistsByPN", "phonenumber", phoneNumber)
	if isExists {
		u.l.Info("User already exists", "id", phoneNumber)
	} else {
		u.l.Info("", "User not find by phonenumber (registration)", phoneNumber)
	}
	return isExists, nil

}

func (u *userService) UserDataForChatHeader(ctx context.Context, id uuid.UUID) (dto.ChatHeader, error) {
	chatHeader, err := u.us.UserDataForChatHeader(ctx, id)
	if err != nil {
		u.l.Error("Error getting user data", "error", err)
		return dto.ChatHeader{}, err
	}
	u.l.Info("Successfully fetched user data", "id", id)
	client := clients.Session.LoadClient(id)
	if client != nil {
		chatHeader.Status = true
	}
	return chatHeader, nil
}

func (u *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := u.us.DeleteUser(ctx, id); err != nil {
		u.l.Error("Error deleting user", "error", err)
		return err
	}
	u.l.Info("Successfully deleting user")
	return nil
}

func DataModify(dtou dto.UserDTO) struct {
	User dto.UserDTO
	ID   uuid.UUID
	Time time.Time
} {
	id := utils.GenerationUUID()
	date := time.Now()
	hash := utils.HashToPassword(dtou.Password)
	userCrNew := dto.UserDTO{Name: dtou.Name, Password: string(hash), PhoneNumber: dtou.PhoneNumber}
	return struct {
		User dto.UserDTO
		ID   uuid.UUID
		Time time.Time
	}{User: userCrNew, ID: id, Time: date}
}
