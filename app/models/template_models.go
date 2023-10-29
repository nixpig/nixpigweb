package models

type Template struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name_" validate:"max=50"`
	Tmpl string `json:"tmpl_" validate:"max=255"`
}
