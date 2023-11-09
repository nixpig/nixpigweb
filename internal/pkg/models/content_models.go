package models

import "time"

type Content struct {
	Id        string    `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"required,max=255"`
	Subtitle  string    `json:"subtitle" validate:"required,max=255"`
	Slug      string    `json:"slug" validate:"required,max=255"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
	Type      string    `json:"type" validate:"required,oneof=post page"`
}
