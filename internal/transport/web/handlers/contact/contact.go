package contact

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (ch *contactHandler) AllContacts(ctx context.Context, c *gin.Context) {
	connew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	userid, exists := c.Get("user_id")
	if !exists {
		ch.l.Error("userid not found in context")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	useriduuid, err := uuid.Parse(userid.(string))
	if err != nil {
		ch.l.Error("error parsing user id", "error", err, "userid", userid)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	contacts, err := ch.csv.AllContacts(connew, useriduuid)
	if err != nil {
		ch.l.Error("error while getting all contacts", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ch.l.Info("successfully got all contacts")
	c.JSON(http.StatusOK, gin.H{"Contacts": contacts})
}

func (ch *contactHandler) IsMyContact(ctx context.Context, c *gin.Context) {
	connew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	id, exists := c.Get("user_id")
	if !exists {
		ch.l.Error("user id not found in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found in context"})
		return
	}
	idu, err := uuid.Parse(id.(string))
	if err != nil {
		ch.l.Error("error parsing user id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var object struct {
		//idone uuid.UUID
		ContactId uuid.UUID
	}

	if err := c.ShouldBindJSON(&object); err != nil {
		ch.l.Error("error while binding json", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contactId := object.ContactId

	isExists, err := ch.csv.IsMyContact(connew, idu, contactId)
	if err != nil {
		ch.l.Error("error while checking if contact is my", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ch.l.Info("successfully checked mycontact", "contactId", contactId)

	c.JSON(http.StatusOK, gin.H{"My": isExists})
}

func (ch *contactHandler) AddContact(ctx context.Context, c *gin.Context) {
	connew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	id, exists := c.Get("user_id")
	if !exists {
		ch.l.Error("user id not found in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found in context"})
		return
	}
	idu, err := uuid.Parse(id.(string))
	if err != nil {
		ch.l.Error("error parsing user id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var contactdto dto.ContactDTOClient

	if err := c.ShouldBindJSON(&contactdto); err != nil {
		ch.l.Error("error while binding json", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	phonenumber := contactdto.PhoneNumber

	contactid, err := ch.us.UserIdByPhoneNumber(connew, phonenumber)
	if err != nil {
		ch.l.Error("error getting user id by phonenumber", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if contactid == idu {
		ch.l.Error("You cannot add a contact to your user")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "You cannot add a contact to your user"})
		return
	}

	contact := dto.ToDomainContact(idu, contactid)
	if err := ch.csv.AddContact(connew, contact.UserId, contact.Contact); err != nil {
		ch.l.Error("error while adding contact", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ch.l.Info("successfully add contact", "contactid", contact.Contact)
	c.JSON(http.StatusCreated, gin.H{"Contact": contact.Contact})
}

func (ch *contactHandler) DeleteContact(ctx context.Context, c *gin.Context) {
	connew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	id, exists := c.Get("user_id")
	if !exists {
		ch.l.Error("user id not found in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found in context"})
		return
	}

	idu, err := uuid.Parse(id.(string))
	if err != nil {
		ch.l.Error("error parsing user id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contactId := c.Param("id")

	contactIdu := uuid.MustParse(contactId)

	if err := ch.csv.DeleteContact(connew, idu, contactIdu); err != nil {
		ch.l.Error("error while deleting contact", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"Contact": contactId})

}
