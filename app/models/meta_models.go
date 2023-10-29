package models

type Meta struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name_" validate:"max=50"`
	Value string `json:"value_" validate:"max=255"`
}
