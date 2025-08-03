package group

import (
	"context"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain/clients"
	"github.com/MaxKudIT/messkudi/internal/domain/groups"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (gsv *groupService) GroupById(ctx context.Context, id uuid.UUID) (dto.GroupDTODetailsServer, error) {
	group, err := gsv.gs.GroupById(ctx, id)
	if err != nil {
		gsv.l.Error("Error getting group by id %s", id)
		return dto.GroupDTODetailsServer{}, err
	}
	gsv.l.Info("Successfully got group by id %s", id)
	return group, nil
}
func (gsv *groupService) CreateGroup(ctx context.Context, group groups.Group, ownerid uuid.UUID, ids []uuid.UUID) error {
	if err := gsv.gs.SaveGroup(ctx, group, ownerid); err != nil {
		gsv.l.Error("Error creating group %s", group)
		return err
	}

	go func() {
		Session := clients.GetSession()
		for _, id := range ids {

			client := Session.LoadClient(id)
			if client == nil {
				continue
			}
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Заходи в группу, id: %s", group.Id)))
			if err != nil {
				gsv.l.Error("Error writing", "error", err)
				return
			}
			gsv.l.Info("Success writing")
		}
	}()
	gsv.l.Info("Successfully created group %s", group)
	return nil

}
func (gsv *groupService) JoinGroup(ctx context.Context, ug groups.UsersGroups) error {
	if err := gsv.gs.JoinGroup(ctx, ug); err != nil {
		gsv.l.Error("Error joining group %s", ug)
		return err
	}
	gsv.l.Info("Successfully joined group %s", ug)
	return nil

}
func (gsv *groupService) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	if err := gsv.gs.DeleteGroup(ctx, id); err != nil {
		gsv.l.Error("Error deleting group %s", id)
		return err
	}
	gsv.l.Info("Successfully deleted group %s", id)
	return nil
}
