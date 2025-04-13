package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (us *userStorage) UserById(ctx context.Context, id uuid.UUID) (dto.UserDTO, error) {
	userp := dto.UserDTO{}
	const GET_USER_QUERY = "SELECT name, lastname, password, phonenumber FROM users WHERE id = $1"
	if err := us.db.QueryRowContext(ctx, GET_USER_QUERY, id).Scan(&userp.Name, &userp.LastName, &userp.Password, &userp.PhoneNumber); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			us.l.Error("user not found", "error", err)
			return dto.UserDTO{}, err
		case errors.Is(err, context.Canceled):
			us.l.Warn("Query cancelled", "error", err)
			return dto.UserDTO{}, err
		case errors.Is(err, context.DeadlineExceeded):
			us.l.Warn("Query timed out", "error", err)
			return dto.UserDTO{}, err
		default:
			us.l.Error("Query failed", "error", err)
			return dto.UserDTO{}, err
		}
	}
	us.l.Info("Successfully got user", "id", id)
	return userp, nil
}

func (us *userStorage) SaveUser(ctx context.Context, userp domain.User) error {

	const CREATE_USER_QUERY = "INSERT INTO users (name, lastname, password, phonenumber, id, createdat, updatedat, refreshtoken, expiredat) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	if _, err := us.db.ExecContext(ctx, CREATE_USER_QUERY, userp.Name, userp.LastName, userp.Password, userp.PhoneNumber, userp.Id, userp.CreatedAt, userp.UpdatedAt, userp.Token.RefreshToken, userp.ExpiredAt); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			us.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			us.l.Warn("Query timed out", "error", err)
			return err
		default:
			us.l.Error("Query failed", "error", err)
			return err
		}
	}
	us.l.Info("Successfully created user", "id", userp.Id)
	return nil
}

func (us *userStorage) DeleteUser(ctx context.Context, id uuid.UUID) error {
	const DELETE_USER_QUERY = "DELETE FROM users WHERE id = $1"

	if _, err := us.db.ExecContext(ctx, DELETE_USER_QUERY, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			us.l.Error("user not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			us.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			us.l.Warn("Query timed out", "error", err)
			return err
		default:
			us.l.Error("Query failed", "error", err)
			return err
		}
	}
	us.l.Info("Successfully deleted user", "id", id)
	return nil
}

func (us *userStorage) UserByPhoneNumber(ctx context.Context, phoneNumber string) (dto.UserDTO, error) {
	userp := dto.UserDTO{PhoneNumber: phoneNumber}
	const GET_USER_QUERY = "SELECT name, lastname, password FROM users WHERE phonenumber = $1"
	if err := us.db.QueryRowContext(ctx, GET_USER_QUERY, phoneNumber).Scan(&userp.Name, &userp.LastName, &userp.Password); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			us.l.Error("user not found", "error", err)
			return dto.UserDTO{}, err
		case errors.Is(err, context.Canceled):
			us.l.Warn("Query cancelled", "error", err)
			return dto.UserDTO{}, err
		case errors.Is(err, context.DeadlineExceeded):
			us.l.Warn("Query timed out", "error", err)
			return dto.UserDTO{}, err
		default:
			us.l.Error("Query failed", "error", err)
			return dto.UserDTO{}, err
		}
	}
	us.l.Info("Successfully got user", "phonenumber", phoneNumber)
	return userp, nil
}
func (us *userStorage) UserIsExistsByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error) {
	var isExists bool
	const GET_USER_QUERY = "SELECT EXISTS (SELECT 1 FROM users WHERE phonenumber = $1);"
	if err := us.db.QueryRowContext(ctx, GET_USER_QUERY, phoneNumber).Scan(&isExists); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			us.l.Warn("Query cancelled", "error", err)
			return false, err
		case errors.Is(err, context.DeadlineExceeded):
			us.l.Warn("Query timed out", "error", err)
			return false, err
		default:
			us.l.Error("Query failed", "error", err)
			return false, err
		}
	}
	if isExists {
		us.l.Info("User already exists", "id", phoneNumber)
	} else {
		us.l.Info("", "User not find by phonenumber (registration)", phoneNumber)
	}
	return isExists, nil
}

// func (userRep userStorage) UpdateUser(userUpd user.UserUpdate) {
//
//		const UPDATE_USER_QUERY = "UPDATE users SET name = $1, lastname = $2, password = $3, phonenumber = $4"
//		if _, err := userRep.db.Exec(UPDATE_USER_QUERY, userUpd.Name, userUpd.LastName, userUpd.Password, userUpd.PhoneNumber); err != nil {
//			log.Print(exceptions.UpdateUserExc().Error())
//		}
//	}
