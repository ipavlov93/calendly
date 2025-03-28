package test

import (
	"testing"

	"event-calendar/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// ARRANGE
	expectedUser := CreateTestUser(1)
	// MANDATORY STEP
	//set expected userID = 0
	expectedUser.ID = 0

	// happy flow
	t.Run("should create user and init giving fields", func(t *testing.T) {
		// ACT
		testUser := domain.NewUser(
			expectedUser.UUID,
			expectedUser.FirstName,
			expectedUser.LastName,
			expectedUser.EmailAddress,
			expectedUser.Organization,
			expectedUser.Description,
		)

		// ASSERT
		assert.Equal(t, expectedUser, testUser)
	})
}
