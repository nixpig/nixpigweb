package models

import "time"

type Page struct {
	Id          int       `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"max=255"`
	Body        string    `json:"body"`
	Slug        string    `json:"slug" validate:"required,max=255"`
	Status      string    `json:"status" validate:"required,max=10"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      int       `json:"user_id" validate:"required"`
	CategoryId  int       `json:"category_id" validate:"required"`
}
