package userservice

import (
	"context"

	"event-calendar/internal/domain"
)

type UserRepository interface {
	GetUsersCount(ctx context.Context) (int64, error)
	GetUserByID(ctx context.Context, id int64) (domain.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (int64, error)
}
