package contact

import (
	"context"
	"github.com/gin-gonic/gin"
)

type contacthandler interface {
	AllContacts(ctx context.Context, c *gin.Context)
	IsMyContact(ctx context.Context, c *gin.Context)
	AddContact(ctx context.Context, c *gin.Context)
	DeleteContact(ctx context.Context, c *gin.Context)
}

type contactrouter struct {
	ch contacthandler
}

func New(ch contacthandler) *contactrouter {
	return &contactrouter{ch: ch}
}
