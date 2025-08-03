package contact

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/clients"
	"github.com/MaxKudIT/messkudi/internal/domain/contacts"
	"github.com/google/uuid"
)

func (csv *contactService) AddContact(ctx context.Context, userid uuid.UUID, contactid uuid.UUID) error {
	if err := csv.cs.AddContact(ctx, userid, contactid); err != nil {
		csv.l.Error("contact adding failed", "user", userid, "contact", contactid, "err", err)
		return err
	}
	csv.l.Info("contact added", "user", userid, "contact", contactid)
	return nil
}

func (csv *contactService) IsMyContact(ctx context.Context, userid uuid.UUID, contactid uuid.UUID) (bool, error) {
	isExists, err := csv.cs.IsMyContact(ctx, userid, contactid)
	if err != nil {
		csv.l.Error("contact isMyContact failed", "user", userid, "contact", contactid)
		return false, err
	}
	return isExists, nil
}

func (csv *contactService) AllContacts(ctx context.Context, userid uuid.UUID) ([]contacts.ContactPreview, error) {
	contacts, err := csv.cs.AllContacts(ctx, userid)

	for _, contact := range contacts {
		if id := clients.Session.LoadClient(contact.UserId); id != nil {
			contact.Status = true
		}
	}

	if err != nil {
		csv.l.Error("contact list failed", "user", userid, "err", err)
		return nil, err
	}
	csv.l.Info("contact list", "user", userid)
	return contacts, nil
}

func (csv *contactService) DeleteContact(ctx context.Context, userid uuid.UUID, contactid uuid.UUID) error {
	if err := csv.cs.DeleteContact(ctx, userid, contactid); err != nil {
		csv.l.Error("contact deleting failed", "user", userid, "contact", contactid, "err", err)
		return err
	}
	csv.l.Info("contact deleted", "user", userid, "contact", contactid)
	return nil
}
