package models

type User struct {
	Id           string
	Username     string
	PasswordHash string `json:"-"`
	Roles        []string
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Roles    string `json:"roles" binding:"required"`
}
