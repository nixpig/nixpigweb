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

type ChangePassword struct {
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
