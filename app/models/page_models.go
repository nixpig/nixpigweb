package models

import "time"

type Page struct {
	Id          int       `json:"id" validate:"required"`
	Title       string    `json:"title_" validate:"max=255"`
	Body        string    `json:"body_"`
	Slug        string    `json:"slug_" validate:"required,max=255"`
	Status      string    `json:"status_" validate:"required,max=10"`
	CreatedAt   time.Time `json:"created_at_"`
	PublishedAt time.Time `json:"published_at_"`
	UpdatedAt   time.Time `json:"updated_at_"`
	UserId      int       `json:"user_id_" validate:"required"`
	CategoryId  int       `json:"category_id" validate:"required"`
}
