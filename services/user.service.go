package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AkifhanIlgaz/key-aero-api/models"
)

type UserService struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserService(ctx context.Context, db *sql.DB) *UserService {
	return &UserService{
		db:  db,
		ctx: ctx,
	}
}

func (service *UserService) GetUser(username string) (*models.User, error) {
	user := models.User{
		Username: username,
	}

	row := service.db.QueryRow(`
	SELECT id, password_hash FROM users
		WHERE username=$1
	`, username)

	err := row.Scan(&user.Id, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	return &user, nil
}
