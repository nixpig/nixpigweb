package models

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Session struct {
	Id        int
	Token     string
	ExpiresAt int64 // unix timestamp
	IssuedAt  int64 // unix timestamp
	UserId    int
}
