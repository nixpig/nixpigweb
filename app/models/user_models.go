package models

import (
	"time"
)

type User struct {
	Id           int       `json:"id" validate:"required,gte=0"`
	Username     string    `json:"username" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password,omitempty,min=8" validate:"required"`
	IsAdmin      bool      `json:"is_admin" validate:"required"`
	RegisteredAt time.Time `json:"registered_at"`
}

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty,min=8" validate:"required"`
}
