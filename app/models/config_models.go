package models

type Config struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"max=50"`
	Value string `json:"value" validate:"max=50"`
}

type NewConfig struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}
