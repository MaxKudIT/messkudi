package group

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/groups"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"log/slog"
)

type groupStorage interface {
	GroupById(ctx context.Context, id uuid.UUID) (dto.GroupDTODetailsServer, error)
	SaveGroup(ctx context.Context, group groups.Group, ownerid uuid.UUID) error
	JoinGroup(ctx context.Context, ug groups.UsersGroups) error
	DeleteGroup(ctx context.Context, id uuid.UUID) error
}

type groupService struct {
	gs groupStorage
	l  *slog.Logger
}

func New(gs groupStorage, l *slog.Logger) *groupService {
	return &groupService{gs: gs, l: l}
}
