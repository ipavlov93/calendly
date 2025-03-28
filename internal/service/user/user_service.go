package userservice

import (
	"context"

	"event-calendar/internal/domain"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository,
	}
}

func (s UserService) GetUserByID(ctx context.Context, id int64) (user domain.User, err error) {
	user, err = s.repository.GetUserByID(ctx, id)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return user, err
	}

	return user, nil
}

func (s UserService) GetUserByUUID(ctx context.Context, uuid string) (user domain.User, err error) {
	user, err = s.repository.GetUserByUUID(ctx, uuid)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return user, err
	}

	return user, nil
}

func (s UserService) CreateUser(ctx context.Context, user domain.User) (int64, error) {
	userID, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		//if errors.Is(err, repository.ErrDuplicate) {
		//return customError with status code BadRequest
		//}
		return 0, err
	}

	return userID, nil
}
