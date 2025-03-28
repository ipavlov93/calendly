package domain

import (
	"fmt"
	"strings"
)

type Participant struct {
	ID int64
	// fields set by participant himself or auth provider
	FirstName      string
	LastName       string
	ContactEmail   string
	User           User
	Organization   string
	Description    string
	AvatarFileName string
}

func NewParticipant(
	firstName string,
	lastName string,
	emailAddress string,
	organization string,
	description string,
	avatarFileName string,
	userID ...int64,
) Participant {
	uID := int64(0)
	if len(userID) > 0 {
		uID = userID[0]
	}

	return Participant{
		FirstName:    firstName,
		LastName:     lastName,
		ContactEmail: emailAddress,
		User: User{
			ID: uID,
		},
		Organization:   organization,
		Description:    description,
		AvatarFileName: avatarFileName,
	}
}

func (p Participant) String() string {
	return fmt.Sprintf("ID:%d, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p Participant) Equals(participant *Participant) bool {
	if p.ID != participant.ID {
		return false
	}
	if !strings.EqualFold(p.FirstName, participant.FirstName) {
		return false
	}
	if !strings.EqualFold(p.LastName, participant.LastName) {
		return false
	}
	return strings.EqualFold(p.ContactEmail, participant.ContactEmail)
}
