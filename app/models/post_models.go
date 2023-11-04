package models

import "time"

type Post struct {
	Id          int       `json:"id" validate:"required,gte=0"`
	Title       string    `json:"title" validate:"required,max=255"`
	SubTitle    string    `json:"subtitle"`
	Body        string    `json:"body"`
	Slug        string    `json:"slug" validate:"required,max=255"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      int       `json:"user_id" validate:"required,gte=0"`
	CategoryId  int       `json:"category_id"`
}

type NewPost struct {
	UserId     int    `json:"user_id" validate:"required"`
	Title      string `json:"title" validate:"required,max=255"`
	Subtitle   string `json:"subtitle" validate:"max=255"`
	Body       string `json:"body"`
	Status     int    `json:"status"`
	CategoryId int    `json:"category_id"`
}
