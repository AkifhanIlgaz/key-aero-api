package models

type User struct {
	Id           string   `json:"id"`
	Username     string   `json:"username" binding:"required"`
	Name         string   `json:"name" binding:"required"`
	LastName     string   `json:"lastName"`
	Password     string   `json:"password,omitempty" binding:"required"`
	PasswordHash string   `json:"-"`
	Roles        []string `json:"roles" binding:"required"`
	Email        string   `json:"email" binding:"required,email"`
	Phone        string   `json:"phone" binding:"required"`
	Department   string   `json:"department" binding:"required"`
}

type SearchUserInput struct {
	Username   string `form:"username"`
	Roles      string `form:"roles"`
	Email      string `form:"email"`
	Phone      string `form:"phone"`
	Department string `form:"department"`
}
