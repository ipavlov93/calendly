package test

import (
	"testing"

	"event-calendar/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewParticipant(t *testing.T) {
	// ARRANGE
	participant1 := CreateTestParticipant(1)

	// happy flow
	t.Run("should create participant and init giving fields (with user ID)", func(t *testing.T) {
		// ACT
		participant2 := domain.NewParticipant(
			participant1.FirstName,
			participant1.LastName,
			participant1.ContactEmail,
			participant1.Organization,
			participant1.Description,
			participant1.AvatarFileName,
		)

		// ASSERT
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.ContactEmail, participant1.ContactEmail)
		assert.Equal(t, participant2.Organization, participant1.Organization)
		assert.Equal(t, participant2.Description, participant1.Description)
		assert.Equal(t, participant2.AvatarFileName, participant1.AvatarFileName)
	})
}
