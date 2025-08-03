package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/group_message_dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (gmh *groupMessageHandler) MessageById(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id := c.Param("id")
	idu, err := uuid.Parse(id)
	if err != nil {
		gmh.l.Error("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	message, err := gmh.gmsv.MessageById(ctxnew, idu)
	if err != nil {
		gmh.l.Error("Error getting group message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gmh.l.Info("Successfully getting group message")
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (gmh *groupMessageHandler) CreateMessage(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, exists := c.Get("user_id")
	if !exists {
		gmh.l.Error("user id not found in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found in context"})
		return
	}
	idu, err := uuid.Parse(id.(string))
	if err != nil {
		gmh.l.Error("error parsing user id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var messagedto group_message_dto.GroupMessageDTOClient

	if err := c.ShouldBindJSON(&messagedto); err != nil {
		gmh.l.Error("Error parsing message", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageid := uuid.New()

	message := group_message_dto.ToDomain(messageid, idu, time.Now(), time.Now(), messagedto)

	if err := gmh.gmsv.CreateMessage(ctxnew, message); err != nil {
		gmh.l.Error("Error creating group message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	gmh.l.Info("Successfully created group message")
	c.JSON(http.StatusCreated, gin.H{"messageid": message.Id})
}

func (gmh *groupMessageHandler) DeleteMessage(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id := c.Param("id")
	idu, err := uuid.Parse(id)
	if err != nil {
		gmh.l.Error("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := gmh.gmsv.DeleteMessage(ctxnew, idu); err != nil {
		gmh.l.Error("Error deleting group message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gmh.l.Info("Successfully deleting group message")
	c.JSON(http.StatusOK, gin.H{"message ": idu})
}
