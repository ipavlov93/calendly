package test

import (
	"testing"

	"event-calendar/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewParticipant(t *testing.T) {
	// ARRANGE
	expectedParticipant := CreateTestParticipant(1)
	// MANDATORY STEP
	//set expected participantID = 0
	expectedParticipant.ID = 0

	expectedUserID := int64(1)
	expectedParticipant.User.ID = expectedUserID

	// happy flow
	t.Run("should create participant and init giving fields (with user ID)", func(t *testing.T) {
		// ACT
		testParticipant := domain.NewParticipant(
			expectedParticipant.FirstName,
			expectedParticipant.LastName,
			expectedParticipant.ContactEmail,
			expectedParticipant.Organization,
			expectedParticipant.Description,
			expectedParticipant.AvatarFileName,
			expectedUserID,
		)

		// ASSERT
		assert.Equal(t, expectedParticipant, testParticipant)
	})
}
