package models

type Status struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
