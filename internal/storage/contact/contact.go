package contact

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaxKudIT/messkudi/internal/domain/contacts"
	"github.com/google/uuid"
	"time"
)

//func (cs *contactStorage) ContactById(ctx context.Context, id uuid.UUID) (dto.UserDTO, error) {
//	userp := dto.UserDTO{}
//	const GET_CONTACT_QUERY = "SELECT name, lastname, password, phonenumber FROM users WHERE id = $1"
//	if err := cs.db.QueryRowContext(ctx, GET_CONTACT_QUERY, id).Scan(&userp.Name, &userp.LastName, &userp.Password, &userp.PhoneNumber); err != nil {
//		switch {
//		case errors.Is(err, sql.ErrNoRows):
//			cs.l.Error("contact not found", "error", err)
//			return dto.UserDTO{}, err
//		case errors.Is(err, context.Canceled):
//			cs.l.Warn("Query cancelled", "error", err)
//			return dto.UserDTO{}, err
//		case errors.Is(err, context.DeadlineExceeded):
//			cs.l.Warn("Query timed out", "error", err)
//			return dto.UserDTO{}, err
//		default:
//			cs.l.Error("Query failed", "error", err)
//			return dto.UserDTO{}, err
//		}
//	}
//	cs.l.Info("Successfully got contact", "id", id)
//	return userp, nil
//}

func (cs *contactStorage) AddContact(ctx context.Context, userid uuid.UUID, contactid uuid.UUID) error {
	now := time.Now()
	const ADD_CONTACT_QUERY = "INSERT INTO users_contacts (user_id, contact_id, adding_at) VALUES ($1, $2, $3)"

	if _, err := cs.db.ExecContext(ctx, ADD_CONTACT_QUERY, userid, contactid, now); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return err
		default:
			cs.l.Error("Query failed", "error", err)
			return err
		}
	}
	cs.l.Info("Successfully add contact", "id", contactid)
	return nil
}

func (cs *contactStorage) IsMyContact(ctx context.Context, id uuid.UUID, contactId uuid.UUID) (bool, error) {
	var isExists bool
	const ConctactIsExists = "SELECT EXISTS (SELECT 1 FROM users_contacts WHERE user_id = $1 AND contact_id = $2);"
	if err := cs.db.QueryRowContext(ctx, ConctactIsExists, id, contactId).Scan(&isExists); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return false, err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return false, err
		default:
			cs.l.Error("Query failed", "error", err)
			return false, err
		}
	}
	if isExists {
		cs.l.Info("Is my contact", "id", contactId)
	} else {
		cs.l.Info("", "Not my contact", contactId)
	}
	return isExists, nil
}

func (cs *contactStorage) AllContacts(ctx context.Context, userid uuid.UUID) ([]contacts.ContactPreview, error) {
	var cts contacts.Contacts
	cts.Ctcs = make([]contacts.ContactPreview, 0)
	const AllContactQuery = "SELECT id, name, color from users u inner join users_contacts uc on u.id = uc.contact_id WHERE uc.user_id = $1"

	rows, err := cs.db.QueryContext(ctx, AllContactQuery, userid)
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return nil, err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return nil, err
		default:
			cs.l.Error("Query failed", "error", err)
			return nil, err
		}
	}
	for rows.Next() {
		var currentObject contacts.ContactPreview
		if err := rows.Scan(&currentObject.UserId, &currentObject.Name, &currentObject.Color); err != nil {
			cs.l.Error("Scan failed", "error", err)
			return nil, err
		}
		cts.Ctcs = append(cts.Ctcs, currentObject)
	}
	cs.l.Info("Successfully getting contacts", "userid", userid)
	return cts.Ctcs, nil
}

func (cs *contactStorage) DeleteContact(ctx context.Context, id uuid.UUID, contactId uuid.UUID) error {
	const DeleteContactQuery = "DELETE FROM users_contacts WHERE user_id = $1 AND contact_id = $2"

	if _, err := cs.db.ExecContext(ctx, DeleteContactQuery, id, contactId); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			cs.l.Error("contact not found", "error", err)
			return err
		case errors.Is(err, context.Canceled):
			cs.l.Warn("Query cancelled", "error", err)
			return err
		case errors.Is(err, context.DeadlineExceeded):
			cs.l.Warn("Query timed out", "error", err)
			return err
		default:
			cs.l.Error("Query failed", "error", err)
			return err
		}
	}
	cs.l.Info("Successfully deleted contact", "id", id)
	return nil
}
