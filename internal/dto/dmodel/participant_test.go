package dmodel_test

import (
	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParticipant(t *testing.T) {
	// ARRANGE
	participant1 := createDummyParticipant(1)

	// happy flow
	t.Run("should create participant and init giving fields (no user ID)", func(t *testing.T) {
		// ACT
		participant2 := dmodel.NewParticipant(
			participant1.FirstName,
			participant1.LastName,
			participant1.ContactEmail,
		)

		// ASSERT
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.ContactEmail, participant1.ContactEmail)
	})
	// happy flow
	t.Run("should create participant and init giving fields (with user ID)", func(t *testing.T) {
		// ARRANGE
		user1 := domain.createDummyUser(1)

		// ACT
		participant2 := dmodel.NewParticipant(
			participant1.FirstName,
			participant1.LastName,
			participant1.ContactEmail,
			user1.ID,
		)

		// ASSERT
		assert.Equal(t, participant2.UserID, user1.ID)
		assert.Equal(t, participant2.FirstName, participant1.FirstName)
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.ContactEmail, participant1.ContactEmail)
	})
}

// fill all the string fields with given value
func createDummyParticipant(id int) *dmodel.Participant {
	return &dmodel.Participant{
		ID:           int64(id),
		FirstName:    fmt.Sprintf("FirstName %d", id),
		LastName:     fmt.Sprintf("LastName %d", id),
		ContactEmail: fmt.Sprintf("%d@test.com", id),
	}
}
