package group

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/groups"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"log/slog"
)

type groupService interface {
	GroupById(ctx context.Context, id uuid.UUID) (dto.GroupDTODetailsServer, error)
	CreateGroup(ctx context.Context, group groups.Group, ownerid uuid.UUID, ids []uuid.UUID) error
	JoinGroup(ctx context.Context, ug groups.UsersGroups) error
	DeleteGroup(ctx context.Context, id uuid.UUID) error
}

type groupHandler struct {
	gsv groupService
	l   *slog.Logger
}

func New(gsv groupService, l *slog.Logger) *groupHandler {
	return &groupHandler{gsv: gsv, l: l}
}
