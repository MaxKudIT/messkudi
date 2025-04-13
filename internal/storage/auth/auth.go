package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
)

func (as *authStorage) UserAuthData(ctx context.Context, userCr dto.UserCredentials) (domain.AuthData, error) {
	var id uuid.UUID
	var name string
	var hashhex string
	var rt string
	const GET_USERPASS_QUERY = "SELECT name, password, id, refreshtoken FROM users WHERE phonenumber = $1"
	if err := as.db.QueryRowContext(ctx, GET_USERPASS_QUERY, userCr.PhoneNumber).Scan(&name, &hashhex, &id, &rt); err != nil {
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
	return domain.AuthData{Id: id, Password: hashhex, Name: name, Token: domain.Token{RefreshToken: rt}}, nil

}

func (as *authStorage) AccessTokenUpdate(ctx context.Context, idparam uuid.UUID) (domain.AccessTokenUpdateData, error) {
	var rt string
	const GET_TOKEN = "SELECT refreshtoken FROM users WHERE id = $1 AND expiredat > NOW()"
	if err := as.db.QueryRowContext(ctx, GET_TOKEN, idparam).Scan(&rt); err != nil {
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
