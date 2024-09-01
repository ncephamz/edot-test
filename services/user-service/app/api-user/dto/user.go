package dto

import (
	model "user-service/models/user"

	"github.com/google/uuid"
)

type (
	UserResp struct {
		Id string `json:"id" example:"7d51e482-7abd-4eef-aefa-1959a60c2e03"`
	}

	UserDTO struct {
		PhoneNumber  string `json:"phone_numer" example:"081320080972"`
		Email        string `json:"email" example:"ncephamz@gmail.com"`
		Password     string `json:"password" example:"12345678"`
		Name         string `json:"name" example:"Encep Hamzah F R" validate:"required"`
		PhotoProfile string `json:"photo_profile"`
	}

	LoginDTO struct {
		Username string `json:"username" validate:"required" example:"3210808710982738"`
		Password string `json:"password" validate:"required" example:"12345678"`
	}

	LoginResp struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiredIn    int64  `json:"expired_in"`
	}

	RefreshTokenDTO struct {
		AccessToken  string
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
)

func (ud UserDTO) ToEntity(id string) model.UserEntity {
	var user = model.UserEntity{
		UserId:       uuid.New().String(),
		PhoneNumber:  ud.PhoneNumber,
		Email:        ud.Email,
		Password:     ud.Password,
		Name:         ud.Name,
		PhotoProfile: ud.PhotoProfile,
	}

	if id != "" {
		user.UserId = id
	}

	return user
}
