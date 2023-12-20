package services

import (
	"context"
	"database/sql"

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
	return nil, nil
}
