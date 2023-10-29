package models

type Category struct {
	Id         int    `json:"id" validate:"required"`
	Name       string `json:"name_" validate:"max=50"`
	TemplateId int    `json:"template_id_"`
}