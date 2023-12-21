package models

type User struct {
	Id           string
	Username     string
	PasswordHash string `json:"-"`
	Roles        []string
}

type AddUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
}
