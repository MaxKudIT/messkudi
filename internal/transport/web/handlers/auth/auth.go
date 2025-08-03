package auth

import (
	"context"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/utils"
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
	utils.SetTokensAndIdHTTPOnlyCookie(c, authdata.Token.AccessToken, authdata.Token.RefreshToken, authdata.Id)

	//clientid := uuid.New()
	//val := rand.Int()
	////deviceid := utils.GenerationDeviceHash(string(val), string(val-51)) //c.ClientIP(), c.Request.UserAgent()
	////expires := time.Now().Add(24 * 30 * time.Hour)
	////session := session2.Session{authdata.Id, clientid, deviceid, expires, true}
	////if err := ah.ss.CreateSession(ctx, session); err != nil {
	////	ah.l.Error("Error while creating session", "error", err)
	////	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	////	return
	////}

	c.JSON(http.StatusOK, gin.H{"Tokens": authdata.Token, "Id": authdata.Id, "Name": authdata.Name, "PhoneNumber": authdata.PhoneNumber})
}

func (ah *authHandler) AccessTokenUpdate(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "refresh_token is expired"})
	}
	id, err := c.Cookie("user_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "user_id is expired"})
	}

	data, err := ah.as.AccessTokenUpdate(ctxnew, dto.UpdateTokenDTO{Rt: dto.RefreshTokenDTO{RefreshToken: rt}, Id: id})
	if err != nil {
		ah.l.Error("Error while updating access token", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ah.l.Info("Successfully updated access token", data)
	utils.SetAccessTokenHTTPOnlyCookie(c, data.AccessToken)
	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (ah *authHandler) IsAuth(ctx context.Context, c *gin.Context) {

	_, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"IsAuth": false})
	}

	id, err := c.Cookie("user_id")
	fmt.Println(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"IsAuth": false})
	}
	c.JSON(http.StatusOK, gin.H{"IsAuth": true, "Id": id})
}

func (ah *authHandler) Logout(ctx context.Context, c *gin.Context) {

	utils.ClearAllCookies(c)

	ah.l.Info("All cookies cleared")

}
