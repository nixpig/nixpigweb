package user

import (
	"time"
)

type UserModel struct {
	Id           int       `json:"id" validate:"required,gte=0"`
	Username     string    `json:"username" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password,omitempty,min=8" validate:"required"`
	IsAdmin      bool      `json:"is_admin" validate:"required"`
	RegisteredAt time.Time `json:"registered_at"`
}

type NewUserModel struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty,min=8" validate:"required"`
}
