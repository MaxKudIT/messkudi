package user

import (
	"context"
	user2 "github.com/MaxKudIT/messkudi/internal/services/user"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (uh *userhandler) UserById(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	user, err := uh.us.UserById(ctxnew, uuid)
	if err != nil {
		uh.l.Error("Error getting user", "id", id, "err", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	uh.l.Info("Successfully got user", "id", id)
	c.JSON(200, gin.H{"Data": user})
}

func (uh *userhandler) CreateUser(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var userdt dto.UserDTO

	if err := c.ShouldBindJSON(&userdt); err != nil {
		uh.l.Error("Error creating user: data not valid")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := user2.DataModify(userdt)
	userp := dto.ToDomainWithoutRefresh(data.ID, data.Time, data.Time, data.User)
	user, err := uh.us.CreateUser(ctxnew, userp)
	if err != nil {
		uh.l.Error("Error creating user", "id", data.ID, "err", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uh.l.Info("Successfully created user", "id", data.ID)
	c.JSON(201, gin.H{"id:": user.Id})
}

func (uh *userhandler) UserByPhoneNumber(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	phonenumber := c.Param("phonenumber")
	pn, err := utils.ValidatePhone(phonenumber, "RU")
	if err != nil {
		uh.l.Error("Error validating phonenumber", "phone", phonenumber, "err", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uh.us.UserByPhoneNumber(ctxnew, pn)
	if err != nil {
		uh.l.Error("Error getting user", "phonenumber", pn, "err", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uh.l.Info("Successfully got user", "phonenumber", pn)
	c.JSON(200, gin.H{"Data:": user})
}

func (uh *userhandler) UserIsExistsByPhoneNumber(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	phonenumber := c.Param("phonenumber")
	isExists, err := uh.us.UserIsExistsByPhoneNumber(ctxnew, phonenumber)
	if err != nil {
		uh.l.Error("Error getting result isExistsByPN", "phonenumber", phonenumber)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	}
	if isExists {
		c.JSON(200, gin.H{"result": isExists, "message": "Данный пользователь уже существует!", "type": "exists"})
		return
	}
	c.JSON(200, gin.H{"result": isExists, "type": "notexists"})
}

//	func (ush UserHandler) UpdateUser(c *gin.Context) {
//		var user user.UserUpdate
//
//		if err := c.ShouldBindJSON(&user); err != nil {
//			log.Print(exceptions.CreateUserExc().Error())
//		}
//		ush.userserv.UpdateUserService(user)
//	}

func (uh *userhandler) DeleteUser(ctx context.Context, c *gin.Context) {

	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	if err := uh.us.DeleteUser(ctxnew, uuid); err != nil {
		uh.l.Error("Error deleting user", "id", id, "err", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}
	utils.ClearAllCookies(c)
	uh.l.Info("Successfully deleted user", "id", id)
}
