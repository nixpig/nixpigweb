package models

import (
	"time"
)

type User struct {
	Id           int       `json:"id" validate:"required,gte=0"`
	Username     string    `json:"username" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	IsAdmin      bool      `json:"is_admin" validate:"required"`
	Password     string    `json:"password,omitempty,min=8" validate:"required"`
	RegisteredAt time.Time `json:"registered_at"`
	LastLogin    time.Time `json:"last_login"`
	Role         string    `json:"role"`
	Profile      string    `json:"profile"`
}

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty,min=8" validate:"required"`
}
