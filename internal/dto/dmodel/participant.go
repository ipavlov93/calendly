package dmodel

import (
	"database/sql"
	"fmt"
	"time"
)

type Participant struct {
	ID int64 `db:"id"`
	// fields set by participant himself or auth provider
	FirstName      string        `db:"first_name"`
	LastName       string        `db:"last_name"`
	ContactEmail   string        `db:"contact_email"`
	UserID         sql.NullInt64 `db:"user_id"`
	Organization   string        `db:"organization"`
	Description    string        `db:"description"`
	AvatarFileName string        `db:"avatar_file_name"`
	CreatedAt      time.Time     `db:"created_at"`
	UpdatedAt      time.Time     `db:"updated_at"`
	// soft delete
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func NewParticipant(
	firstName string,
	lastName string,
	emailAddress string,
	organization string,
	description string,
	avatarFileName string,
	userID ...int64,
) *Participant {
	uID := int64(0)
	if len(userID) > 0 {
		uID = userID[0]
	}

	return &Participant{
		FirstName:    firstName,
		LastName:     lastName,
		ContactEmail: emailAddress,
		UserID: sql.NullInt64{
			Int64: uID,
			Valid: uID > 0,
		},
		Organization:   organization,
		Description:    description,
		AvatarFileName: avatarFileName,
	}
}

func (p Participant) String() string {
	return fmt.Sprintf("ID:%d, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}
