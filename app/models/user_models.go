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
	LastLogin    time.Time `json:"last_login,omitempty"`
	Role         string    `json:"role"`
	Profile      string    `json:"profile"`
}

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password,omitempty,min=8" validate:"required"`
}

type UserMeta struct {
	Id     int `json:"id" validate:"required"`
	UserId int `json:"user_id" validate:"required"`
	MetaId int `json:"meta_id" validate:"required"`
}
