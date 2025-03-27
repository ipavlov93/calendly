package test

import (
	"event-calendar/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// ARRANGE
	user1 := CreateTestUser(1)

	// happy flow
	t.Run("should create user and init giving fields", func(t *testing.T) {
		// ACT
		user2 := domain.NewUser(
			user1.UUID,
			user1.FirstName,
			user1.LastName,
			user1.EmailAddress,
			user1.Organization,
			user1.Description,
		)

		// ASSERT
		assert.Equal(t, user2.UUID, user1.UUID)
		assert.Equal(t, user2.FirstName, user1.FirstName)
		assert.Equal(t, user2.LastName, user1.LastName)
		assert.Equal(t, user2.EmailAddress, user1.EmailAddress)
		assert.Equal(t, user2.Organization, user1.Organization)
		assert.Equal(t, user2.Description, user1.Description)
	})
}
