package groups

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	Id        uuid.UUID
	Title     string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersGroups struct {
	UserId  uuid.UUID
	GroupId uuid.UUID
}

type Roles string

const (
	Member Roles = "member"
	Admin  Roles = "admin"
	Owner  Roles = "owner"
)
