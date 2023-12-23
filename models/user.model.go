package models

type User struct {
	Id           string   `json:"id"`
	Username     string   `json:"username" binding:"required"`
	Password     string   `json:"-" binding:"required"`
	PasswordHash string   `json:"-"`
	Roles        []string `json:"roles" binding:"required"`
	Email        string   `json:"email" binding:"required"`
	Phone        string   `json:"phone" binding:"required"`
	Department   string   `json:"department" binding:"required"`
}
