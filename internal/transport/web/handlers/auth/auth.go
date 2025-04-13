package auth

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (ah *authHandler) UserAuthData(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var usercr dto.UserCredentials
	if err := c.ShouldBindJSON(&usercr); err != nil {
		ah.l.Error("data not valid! (auth)")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authdata, err := ah.as.UserAuthData(ctxnew, usercr)
	if err != nil {
		if err.Error() == "Данного пользователя не существует" {
			ah.l.Error("user not exists", "error", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ah.l.Error("Error while getting authdata", "error", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ah.l.Info("Successfully auth", "id", authdata.Id)
	c.JSON(http.StatusOK, gin.H{"Data": authdata})
}

func (ah *authHandler) AccessTokenUpdate(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var rt dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&rt); err != nil {
		ah.l.Error("data not valid! (auth)")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := ah.as.AccessTokenUpdate(ctxnew, c, &rt)
	if err != nil {
		ah.l.Error("Error while updating access token", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ah.l.Info("Successfully updated access token", data)
	c.JSON(http.StatusOK, gin.H{"Data": data})
}
