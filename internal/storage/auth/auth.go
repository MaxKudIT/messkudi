package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"time"
)

func (as *authStorage) UserAuthData(ctx context.Context, userCr dto.UserCredentials) (domain.AuthData, error) {
	fmt.Println(userCr.PhoneNumber)
	var (
		id      uuid.UUID
		name    string
		hashhex string
	)

	const GetUserPassQuery = "SELECT name, password, id FROM users WHERE phonenumber = $1"
	if err := as.db.QueryRowContext(ctx, GetUserPassQuery, userCr.PhoneNumber).Scan(&name, &hashhex, &id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			as.l.Error("user not found", "error", err)
			return domain.AuthData{}, fmt.Errorf("Данного пользователя не существует")
		case errors.Is(err, context.Canceled):
			as.l.Warn("Query cancelled", "error", err)
			return domain.AuthData{}, fmt.Errorf("query cancelled: %w", err)
		case errors.Is(err, context.DeadlineExceeded):
			as.l.Warn("Query timed out", "error", err)
			return domain.AuthData{}, fmt.Errorf("query timed out: %w", err)
		default:
			as.l.Error("Query failed", "error", err)
			return domain.AuthData{}, fmt.Errorf("query failed: %w", err)
		}
	}
	as.l.Info("successfuly getting userid", userCr.PhoneNumber)
	return domain.AuthData{Id: id, Password: hashhex, Name: name, PhoneNumber: userCr.PhoneNumber}, nil

}

func (as *authStorage) UserUpdateRefreshToken(ctx context.Context, newRt string, expired time.Time, pn string) error {
	const UpdateRTQuery = "UPDATE users SET refreshtoken = $1, expiredat = $2 WHERE phonenumber = $3"
	if _, err := as.db.ExecContext(ctx, UpdateRTQuery, newRt, expired, pn); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			as.l.Error("user not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			as.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			as.l.Warn("Query timed out", "error", err)
			return err
		default:
			as.l.Error("Query failed", "error", err)
			return err
		}
	}
	as.l.Info("Successfully update refresh token")
	return nil
}

func (as *authStorage) AccessTokenUpdate(ctx context.Context, idparam uuid.UUID) (domain.AccessTokenUpdateData, error) {
	var rt string
	const GetTokenQuery = "SELECT refreshtoken FROM users WHERE id = $1 AND expiredat > NOW()"
	if err := as.db.QueryRowContext(ctx, GetTokenQuery, idparam).Scan(&rt); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			as.l.Error("user not found by refresh token or expired", "error", err)
			return domain.AccessTokenUpdateData{}, errors.New("user not exists")
		case errors.Is(err, context.Canceled):
			as.l.Warn("Query cancelled", "error", err)
			return domain.AccessTokenUpdateData{}, fmt.Errorf("query cancelled: %w", err)
		case errors.Is(err, context.DeadlineExceeded):
			as.l.Warn("Query timed out", "error", err)
			return domain.AccessTokenUpdateData{}, fmt.Errorf("query timed out: %w", err)
		default:
			as.l.Error("Query failed", "error", err)
			return domain.AccessTokenUpdateData{}, fmt.Errorf("query failed: %w", err)
		}
	}
	as.l.Info("successfuly getting data by refresh token", rt)
	return domain.AccessTokenUpdateData{Id: idparam, RefreshToken: rt}, nil
}
