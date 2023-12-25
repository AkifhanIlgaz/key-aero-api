package models

type User struct {
	Id           string   `json:"id"`
	Username     string   `json:"username" binding:"required"`
	Password     string   `json:"-" binding:"required"`
	PasswordHash string   `json:"-"`
	Roles        []string `json:"roles" binding:"required"`
	Email        string   `json:"email" binding:"required,email"`
	Phone        string   `json:"phone" binding:"required"`
	Department   string   `json:"department" binding:"required"`
}

type UpdateInput struct {
	Id         string   `json:"id"`
	Username   string   `json:"username" binding:"required"`
	Roles      []string `json:"roles" binding:"required"`
	Email      string   `json:"email" binding:"required"`
	Phone      string   `json:"phone" binding:"required"`
	Department string   `json:"department" binding:"required"`
}

type SearchInput struct {
	Username   string   `json:"username,omitempty"`
	Roles      []string `json:"roles,omitempty"`
	Email      string   `json:"email" binding:"email"`
	Phone      string   `json:"phone,omitempty"`
	Department string   `json:"department,omitempty"`
}
