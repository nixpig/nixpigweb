package models

import "time"

type Post struct {
	Id          int       `json:"id" validate:"required,gte=0"`
	Title       string    `json:"title" validate:"required,max=255"`
	Body        string    `json:"body"`
	Views       int       `json:"views" validate:"required,gte=0"`
	Slug        string    `json:"slug" validate:"required,max=255"`
	Published   bool      `json:"published" validate:"required"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      int       `json:"user_id" validate:"required,gte=0"`
}

type NewPost struct {
	UserId int    `json:"user_id" validate:"required,gte=0"`
	Title  string `json:"title" valiate:"required,max=255"`
	Body   string `json:"body"`
}
