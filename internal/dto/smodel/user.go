package smodel

import (
	"fmt"
	"net/mail"
)

type User struct {
	ID   int64  `json:"id"`
	UUID string `json:"uuid"`
	// fields set by participant himself or auth provider
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	EmailAddress mail.Address `json:"email_address"`
	Organization string       `json:"organization"`
	Description  string       `json:"description"`
}

func (p User) String() string {
	return fmt.Sprintf("ID: %d, FullName: %s %s, UUID: %s, Email: %s", p.ID, p.FirstName, p.LastName, p.UUID, p.EmailAddress)
}
