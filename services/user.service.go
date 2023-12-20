package services

import (
	"context"
	"database/sql"
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
