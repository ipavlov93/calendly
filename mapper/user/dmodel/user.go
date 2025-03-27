package dmodel

import (
	"database/sql"
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
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(user.Description),
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
		Description:  user.Description.String,
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
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(user.Description),
	}
}
