package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id           int       `json:"id" validate:"required,gte=0"`
	Username     string    `json:"username" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	IsAdmin      bool      `json:"is_admin" validate:"required"`
	RegisteredAt time.Time `json:"registered_at"`
}

func (u User) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *User) Scan(value interface{}) error {
	i, ok := value.([]byte)
	if !ok {
		return errors.New("failed to assert type as []byte")
	}

	return json.Unmarshal(i, &u)
}

func (u *User) Validate() (*User, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(u)
	if err != nil {
		return nil, err
	}

	return u, err
}
