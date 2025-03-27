package dmodel

import (
	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	"event-calendar/internal/dto/smodel"
)

func UserToUserDto(user domain.User) dmodel.User {
	return dmodel.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		Organization: user.Organization,
		Description:  user.Description,
	}
}

func UserDtoToUser(user dmodel.User) domain.User {
	return domain.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		Organization: user.Organization,
		Description:  user.Description,
	}
}

// MapDto maps Smodel.User to Dmodel.User
func MapDto(user smodel.User) dmodel.User {
	return dmodel.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress.Address,
		Organization: user.Organization,
		Description:  user.Description,
	}
}
