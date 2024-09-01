package user

import (
	"time"
)

type (
	UserEntity struct {
		UserId       string    `json:"user_id" db:"user_id"`
		PhoneNumber  string    `json:"phone_number" db:"phone_number"`
		Email        string    `json:"email" db:"email"`
		Password     string    `json:"password" db:"password"`
		Name         string    `json:"name" db:"name"`
		PhotoProfile string    `json:"photo_profile" db:"photo_profile"`
		CreatedAt    time.Time `json:"created_at" db:"created_at"`
	}
)
