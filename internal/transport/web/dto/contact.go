package dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain/contacts"
	"github.com/google/uuid"
)

type ContactDTOClient struct {
	PhoneNumber string
}

func ToDomainContact(userid uuid.UUID, contactid uuid.UUID) contacts.Contact {
	return contacts.Contact{
		UserId:  userid,
		Contact: contactid,
	}
}
