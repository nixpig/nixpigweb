package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	IsAdmin  bool   `json:"is_admin"`
}
