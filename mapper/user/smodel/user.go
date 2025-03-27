package smodel

import (
	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	"event-calendar/internal/dto/smodel"
	"fmt"
	"net/mail"
)

func UserToUserDto(user domain.User) smodel.User {
	return smodel.User{
		ID:        user.ID,
		UUID:      user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		EmailAddress: mail.Address{
			Name:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Address: user.EmailAddress,
		},
		Organization: user.Organization,
		Description:  user.Description,
	}
}

func UserDtoToUser(user smodel.User) domain.User {
	return domain.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress.Address,
		Organization: user.Organization,
		Description:  user.Description,
	}
}

// MapDto maps Dmodel.User to Smodel.User
func MapDto(user dmodel.User) smodel.User {
	return smodel.User{
		ID:        user.ID,
		UUID:      user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		EmailAddress: mail.Address{
			Name:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Address: user.EmailAddress,
		},
		Organization: user.Organization,
		Description:  user.Description,
	}
}
