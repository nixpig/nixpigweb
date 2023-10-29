package models

type Meta struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"max=50"`
	Value string `json:"value" validate:"max=255"`
}
