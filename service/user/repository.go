package userservice

import (
	"context"

	"event-calendar/internal/dto/dmodel"
)

type UserRepository interface {
	GetUsersCount(ctx context.Context) (int64, error)
	GetUserByID(ctx context.Context, id int64) (dmodel.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (dmodel.User, error)
	CreateUser(ctx context.Context, user dmodel.User) (int64, error)
}
