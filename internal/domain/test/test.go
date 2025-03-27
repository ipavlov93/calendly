package test

import (
	"event-calendar/internal/domain"
	"fmt"
)

// CreateTestUser fill all fields based on given value
func CreateTestUser(id int) domain.User {
	return domain.User{
		ID:           int64(id),
		UUID:         fmt.Sprint(id),
		FirstName:    fmt.Sprintf("FirstName%d", id),
		LastName:     fmt.Sprintf("LastName%d", id),
		EmailAddress: fmt.Sprintf("%d@test.com", id),
		Organization: fmt.Sprintf("Organization%d", id),
		Description:  fmt.Sprintf("Description%d", id),
	}
}

// CreateTestParticipant fill all fields based on given value
func CreateTestParticipant(id int) domain.Participant {
	return domain.Participant{
		ID:             int64(id),
		FirstName:      fmt.Sprintf("FirstName%d", id),
		LastName:       fmt.Sprintf("LastName%d", id),
		ContactEmail:   fmt.Sprintf("%d@test.com", id),
		Organization:   fmt.Sprintf("Organization%d", id),
		Description:    fmt.Sprintf("Description%d", id),
		AvatarFileName: fmt.Sprintf("AvatarFileName%d", id),
	}
}
