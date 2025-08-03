package contacts

import "github.com/google/uuid"

type Contact struct {
	UserId  uuid.UUID
	Contact uuid.UUID
}

type Contacts struct {
	Ctcs []ContactPreview
}

type ContactPreview struct {
	UserId uuid.UUID
	Status bool
	Color  string
	Name   string
}
