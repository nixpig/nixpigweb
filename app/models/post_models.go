package models

import "time"

type Post struct {
	Id          int       `json:"id" validate:"required,gte=0"`
	Title       string    `json:"title_" validate:"required,max=255"`
	SubTitle    string    `json:"subtitle_"`
	Body        string    `json:"body_"`
	Slug        string    `json:"slug_" validate:"required,max=255"`
	Status      string    `json:"status_"`
	CreatedAt   time.Time `json:"created_at_"`
	PublishedAt time.Time `json:"published_at_"`
	UpdatedAt   time.Time `json:"updated_at_"`
	UserId      int       `json:"user_id_" validate:"required,gte=0"`
	CategoryId  int       `json:"category_id_"`
}

type NewPost struct {
	UserId     int    `json:"user_id_" validate:"required"`
	Title      string `json:"title_" validate:"required,max=255"`
	Subtitle   string `json:"subtitle_" validate:"max=255"`
	Body       string `json:"body_"`
	Status     string `json:"status_"`
	CategoryId int    `json:"category_id_"`
}
