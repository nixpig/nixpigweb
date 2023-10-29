package models

type Template struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"max=50"`
	Tmpl string `json:"tmpl" validate:"max=255"`
}
