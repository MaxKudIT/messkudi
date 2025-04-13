package user

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
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
	c.JSON(200, gin.H{"Data:": user})
}

func (uh *userhandler) CreateUser(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var userdt dto.UserDTO

	if err := c.ShouldBindJSON(&userdt); err != nil {
		uh.l.Error("Error creating user: data not valid")
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	id := utils.GenerationUUID()
	location, err := time.LoadLocation("UTC")
	date := time.Now().In(location).Format("2006-01-02")
	hash := utils.HashToPassword(userdt.Password)
	rt := utils.CreateRefreshToken()
	expired := time.Now().Add(30 * 24 * time.Hour).In(location).Format("2006-01-02")
	userCrNew := dto.UserDTO{Name: userdt.Name, LastName: userdt.LastName, Password: string(hash), PhoneNumber: userdt.PhoneNumber}
	userp := dto.ToDomain(id, date, date, domain.Token{RefreshToken: rt}, expired, userCrNew)
	user, err := uh.us.CreateUser(ctxnew, userp)
	if err != nil {
		uh.l.Error("Error creating user", "id", id, "err", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return

	}
	uh.l.Info("Successfully created user", "id", id)
	c.JSON(201, gin.H{"Data:": user})
}

func (uh *userhandler) UserByPhoneNumber(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	phonenumber := c.Param("phonenumber")

	user, err := uh.us.UserByPhoneNumber(ctxnew, phonenumber)
	if err != nil {
		uh.l.Error("Error getting user", "phonenumber", phonenumber, "err", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uh.l.Info("Successfully got user", "phonenumber", phonenumber)
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
	uh.l.Info("Successfully deleted user", "id", id)
}
