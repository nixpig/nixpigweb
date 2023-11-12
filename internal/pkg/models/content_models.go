package models

import "time"

type Content struct {
	Id        string    `json:"id"`
	Title     string    `json:"title" validate:"required,max=255"`
	Subtitle  string    `json:"subtitle" validate:"required,max=255"`
	Slug      string    `json:"slug"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Type      string    `json:"type" validate:"required,oneof=post page"`
}
