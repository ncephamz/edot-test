package user

import (
	"context"
	modelUser "user-service/models/user"
)

type UserStorer interface {
	CreateOne(ctx context.Context, user modelUser.UserEntity) error
	GetByEmailOrPhoneNumber(ctx context.Context, username string) (modelUser.UserEntity, error)
}
