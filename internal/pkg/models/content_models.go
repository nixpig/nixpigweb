package models

import "time"

type Content struct {
	Id        int       `json:"id"`
	Title     string    `json:"title" validate:"required,max=255"`
	Subtitle  string    `json:"subtitle" validate:"required,max=255"`
	Slug      string    `json:"slug" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Type      string    `json:"type" validate:"required,oneof=post page"`
	UserId    int       `json:"user_id"`
}
