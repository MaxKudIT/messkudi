package auth

import (
	"context"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
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
	authdata.Token.AccessToken = acctoken

	as.l.Info("Successfully got authdata", "id", authdata.Id)
	return authdata, nil
}

func (as *authService) AccessTokenUpdate(ctx context.Context, c *gin.Context, rt *dto.RefreshTokenDTO) (domain.AccessTokenUpdateData, error) {
	user_id, exists := c.Get("user_id")
	if !exists {
		as.l.Error("Userid not found in gin context")
		return domain.AccessTokenUpdateData{}, errors.New("Userid not found in gin context")
	}
	user_id_orig, err := uuid.Parse(user_id.(string))
	if err != nil {
		as.l.Error("Userid not found in gin context")
		return domain.AccessTokenUpdateData{}, errors.New("Userid not found in gin context")
	}
	data, err := as.ast.AccessTokenUpdate(ctx, user_id_orig)
	if err != nil {
		as.l.Error("Error while updating access token", "error", err)
		return domain.AccessTokenUpdateData{}, err
	}
	if !(data.RefreshToken == rt.RefreshToken) {
		as.l.Error("refreshtoken not valid", "error", err)
		return domain.AccessTokenUpdateData{}, err
	}
	as.l.Info("Successfully updated access token", "id", data.Id)

	if err != nil {
		as.l.Error("Invalid user_id type in gin context: expected string")
		return domain.AccessTokenUpdateData{}, errors.New("Invalid user_id type in gin context: expected string")
	}
	acctoken, err := utils.CreateAccessToken(os.Getenv("SECRET_KEY"), user_id_orig)
	data.AccessToken = acctoken
	return data, nil
}
