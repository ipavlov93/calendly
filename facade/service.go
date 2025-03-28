package facade

import (
	"context"

	"event-calendar/internal/domain"
)

//type AuthService interface {
//	VerifyToken(token string) (claims auth.Claims, error)
//}

type UserService interface {
	GetUserByID(ctx context.Context, id int64) (userDto domain.User, err error)
	GetUserByUUID(ctx context.Context, uuid string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (int64, error)
}
