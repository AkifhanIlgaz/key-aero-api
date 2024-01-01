package models

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignOutCredentials struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type RefreshCredentials struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type ChangePasswordCredentials struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
