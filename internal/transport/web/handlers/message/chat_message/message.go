package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cmh *chatMessageHandler) MessageById(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	id := c.Param("id")
	idu, err := uuid.Parse(id)
	if err != nil {
		cmh.l.Error("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	message, err := cmh.cmsv.MessageById(ctxnew, idu)
	if err != nil {
		cmh.l.Error("Error getting chat message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cmh.l.Info("Successfully getting chat message")
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (cmh *chatMessageHandler) CreateMessage(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var messagedto chat_message_dto.ChatMessageDTOClient

	if err := c.ShouldBindJSON(&messagedto); err != nil {
		cmh.l.Error("Error parsing message", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messagedtoP, err := messagedto.UuidParse()
	if err != nil {
		cmh.l.Error("Error parsing message", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := chat_message_dto.ToDomain(time.Now(), time.Now(), nil, messagedtoP)

	if err := cmh.cmsv.CreateMessage(ctxnew, message); err != nil {
		cmh.l.Error("Error creating chat message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	cmh.l.Info("Successfully created chat message")
	c.JSON(http.StatusCreated, gin.H{"messageid": message.Id})
}

func (cmh *chatMessageHandler) AllMessages(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	chatid := c.Param("id")
	chatiduuid, err := uuid.Parse(chatid)
	if err != nil {
		cmh.l.Error("Error parsing chatid", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, err := cmh.cmsv.AllMessages(ctxnew, chatiduuid)
	if err != nil {
		cmh.l.Error("Error getting all messages", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func (cmh *chatMessageHandler) UpdateReadAtMessage(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var object struct {
		Tm time.Time
		Id uuid.UUID
	}

	if err := c.ShouldBindJSON(&object); err != nil {
		cmh.l.Error("Error parsing object updateReadAtMessage", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cmh.cmsv.UpdateReadAtMessage(ctxnew, object.Tm, object.Id); err != nil {
		cmh.l.Error("Error parsing object updateReadAtMessage", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	cmh.l.Info("Successfully updated readat chat message")
	c.JSON(http.StatusCreated, gin.H{"messageid": object.Id})
}

func (cmh *chatMessageHandler) DeleteMessage(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id := c.Param("id")
	idu, err := uuid.Parse(id)
	if err != nil {
		cmh.l.Error("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := cmh.cmsv.DeleteMessage(ctxnew, idu); err != nil {
		cmh.l.Error("Error deleting chat message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cmh.l.Info("Successfully deleting chat message")
	c.JSON(http.StatusOK, gin.H{"message deleting: ": idu})
}
