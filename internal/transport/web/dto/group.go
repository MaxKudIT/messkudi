package dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain/groups"
	"github.com/google/uuid"
	"time"
)

type GroupDTOClient struct {
	Title     string
	AvatarURL string
	Ids       []uuid.UUID
}

type GroupDTODetailsServer struct {
	Title       string
	AvatarURL   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	GroupCount  int
}

func ToDomainGroup(groupid uuid.UUID, createdat time.Time, updatedat time.Time, dto GroupDTOClient) groups.Group {
	return groups.Group{
		Id:        groupid,
		Title:     dto.Title,
		AvatarURL: dto.AvatarURL,
		CreatedAt: createdat,
		UpdatedAt: updatedat,
	}
}

//DESCRIPTION В UPDATE

type UserGroupDTOClient struct {
	GroupId uuid.UUID
}

func ToDomainUserGroup(userid uuid.UUID, dto UserGroupDTOClient) groups.UsersGroups {
	return groups.UsersGroups{
		UserId:  userid,
		GroupId: dto.GroupId,
	}
}
