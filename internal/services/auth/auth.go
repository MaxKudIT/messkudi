package auth

import (
	"context"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/utils"
	"github.com/google/uuid"
	"os"
	"time"
)

func (as *authService) UserAuthData(ctx context.Context, usercr dto.UserCredentials) (domain.AuthData, error) {
	authdata, err := as.ast.UserAuthData(ctx, usercr)
	if err != nil {
		as.l.Error("Error while getting authdata", "error", err)
		return domain.AuthData{}, err
	}
	if err = utils.CompareToHash([]byte(authdata.Password), []byte(usercr.Password)); err != nil {
		as.l.Error("pass not valid", "error", err)
		return domain.AuthData{}, errors.New("Неверный пароль")
	}
	acctoken, err := utils.CreateAccessToken(os.Getenv("SECRET_KEY"), authdata.Id)
	if err != nil {
		as.l.Error("create access token failed")
		return domain.AuthData{}, err
	}
	newRt := utils.CreateRefreshToken()
	hashRt := utils.HashRefreshToken(newRt)
	expired := time.Now().Add(time.Hour * 24 * 30)
	authdata.Token = domain.Token{
		AccessToken:  acctoken,
		RefreshToken: newRt,
	}
	if err := as.ast.UserUpdateRefreshToken(ctx, hashRt, expired, usercr.PhoneNumber); err != nil {
		as.l.Error("update refresh token failed")
		return domain.AuthData{}, err
	}
	as.l.Info("Successfully got authdata", "id", authdata.Id)
	return authdata, nil
}

func (as *authService) AccessTokenUpdate(ctx context.Context, dd dto.UpdateTokenDTO) (domain.AccessTokenUpdateData, error) {
	idp, err := uuid.Parse(dd.Id)
	if err != nil {
		return domain.AccessTokenUpdateData{}, err
	}
	data, err := as.ast.AccessTokenUpdate(ctx, idp)
	if err != nil {
		as.l.Error("Error while updating access token", "error", err)
		return domain.AccessTokenUpdateData{}, err
	}
	if err := utils.CompareRefreshToken([]byte(data.RefreshToken), []byte(dd.Rt.RefreshToken)); err != nil {
		as.l.Error("refreshtoken not valid")
		return domain.AccessTokenUpdateData{}, errors.New("refreshtoken not valid")
	}
	as.l.Info("Successfully updated access token", "id", data.Id)

	acctoken, err := utils.CreateAccessToken(os.Getenv("SECRET_KEY"), idp)
	data.AccessToken = acctoken
	return data, nil
}
