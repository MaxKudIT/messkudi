package group

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/groups"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
)

func (gs *groupStorage) GroupById(ctx context.Context, id uuid.UUID) (dto.GroupDTODetailsServer, error) {
	var group dto.GroupDTODetailsServer
	const GET_GROUP_QUERY = "SELECT title, description, color, createdat, updatedat, group_count from groups where id = $1"
	if err := gs.db.QueryRowContext(ctx, GET_GROUP_QUERY, id).Scan(&group.Title, &group.Description, &group.Color, &group.CreatedAt, &group.UpdatedAt, &group.GroupCount); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			gs.l.Error("group not found", "error", err)
			return dto.GroupDTODetailsServer{}, err
		case errors.Is(err, context.Canceled):
			gs.l.Warn("Query cancelled", "error", err)
			return dto.GroupDTODetailsServer{}, err
		case errors.Is(err, context.DeadlineExceeded):
			gs.l.Warn("Query timed out", "error", err)
			return dto.GroupDTODetailsServer{}, err
		default:
			gs.l.Error("Query failed", "error", err)
			return dto.GroupDTODetailsServer{}, err
		}
	}
	gs.l.Info("Successfully got group", "id", id)
	return group, nil
}

func (gs *groupStorage) SaveGroup(ctx context.Context, group groups.Group, ownerid uuid.UUID) error {
	var role groups.Roles = groups.Owner
	const SAVE_GROUP_QUERY = "BEGIN; INSERT INTO groups (id, title, avatarurl, createdat, updatedat) VALUES ($1, $2, $3, $4, $5); INSERT INTO users_groups (user_id, group_id, role) VALUES ($6, $7, $8); COMMIT;)"
	if _, err := gs.db.ExecContext(ctx, SAVE_GROUP_QUERY, group.Id, group.Title, group.AvatarURL, group.CreatedAt, group.UpdatedAt, ownerid, group.Id, role); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			gs.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			gs.l.Warn("Query timed out", "error", err)
			return err
		default:
			gs.l.Error("Query failed", "error", err)
			return err
		}
	}
	gs.l.Info("Successfully saving group", "id", group.Id)
	return nil
}

func (gs *groupStorage) JoinGroup(ctx context.Context, ug groups.UsersGroups) error {

	const JOIN_GROUP_QUERY = "INSERT INTO users_groups (user_id, group_id) VALUES ($1, $2, $3)"

	if _, err := gs.db.ExecContext(ctx, JOIN_GROUP_QUERY, ug.UserId, ug.GroupId); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			gs.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			gs.l.Warn("Query timed out", "error", err)
			return err
		default:
			gs.l.Error("Query failed", "error", err)
			return err
		}
	}
	gs.l.Info("Successfully join to group", "GroupId", ug.GroupId, "UserId", ug.UserId)
	return nil
}

func (gs *groupStorage) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	const DELETE_GROUP_QUERY = "DELETE FROM groups WHERE id = $1"

	if _, err := gs.db.ExecContext(ctx, DELETE_GROUP_QUERY, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			gs.l.Error("group not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			gs.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			gs.l.Warn("Query timed out", "error", err)
			return err
		default:
			gs.l.Error("Query failed", "error", err)
			return err
		}
	}
	gs.l.Info("Successfully deleted group", "id", id)
	return nil
}
