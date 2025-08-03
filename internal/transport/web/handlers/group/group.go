package group

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func (gh *groupHandler) GroupById(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id := c.Param("id")

	group, err := gh.gsv.GroupById(ctxnew, func() uuid.UUID { idu, _ := uuid.Parse(id); return idu }())
	if err != nil {
		gh.l.Error("Error getting group", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	gh.l.Info("Successfully got group", "group", group)
	c.JSON(http.StatusOK, gin.H{"group": group})

}
func (gh *groupHandler) CreateGroup(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, exists := c.Get("user_id")
	if !exists {
		gh.l.Error("user id not found in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found in context"})
		return
	}
	ownerid, err := uuid.Parse(id.(string))
	if err != nil {
		gh.l.Error("error parsing user id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var gdto dto.GroupDTOClient

	if err := c.ShouldBindJSON(&gdto); err != nil {
		gh.l.Error("error while binding json", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	groupid := uuid.New()
	group := dto.ToDomainGroup(groupid, time.Now(), time.Now(), gdto)

	if err := gh.gsv.CreateGroup(ctxnew, group, ownerid, gdto.Ids); err != nil {
		gh.l.Error("Error creating group", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	gh.l.Info("Successfully created group", "group", group)
	c.JSON(http.StatusOK, gin.H{"groupid": group.Id}) //gggggggggtad
}
func (gh *groupHandler) JoinGroup(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, exists := c.Get("user_id")
	if !exists {
		gh.l.Error("user id not found in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found in context"})
		return
	}
	idu, err := uuid.Parse(id.(string))
	if err != nil {
		gh.l.Error("error parsing user id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var ugdto dto.UserGroupDTOClient
	if err := c.ShouldBindJSON(&ugdto); err != nil {
		gh.l.Error("Error parsing request body", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ug := dto.ToDomainUserGroup(idu, ugdto)

	if err := gh.gsv.JoinGroup(ctxnew, ug); err != nil {
		gh.l.Error("Error joining group", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	gh.l.Info("Successfully joined group", "group", ug)
	c.JSON(http.StatusOK, gin.H{"groupid": ug.GroupId})

}
func (gh *groupHandler) DeleteGroup(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id := c.Param("id")

	if err := gh.gsv.DeleteGroup(ctxnew, func() uuid.UUID { idu, _ := uuid.Parse(id); return idu }()); err != nil {
		gh.l.Error("Error deleting group", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	gh.l.Info("Successfully deleted group", "group", id)
	c.JSON(http.StatusOK, gin.H{"group": nil})
}
