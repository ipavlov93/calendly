package dmodel

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID   int64  `db:"id"`
	UUID string `db:"uuid"`
	// fields set by participant himself or auth provider
	FirstName    string         `db:"first_name"`
	LastName     string         `db:"last_name"`
	EmailAddress string         `db:"email_address"`
	Organization string         `db:"organization"`
	Description  sql.NullString `db:"description"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	// soft delete
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (p User) String() string {
	return fmt.Sprintf("ID: %d, FullName: %s %s, UUID: %s, Email: %s", p.ID, p.FirstName, p.LastName, p.UUID, p.EmailAddress)
}
