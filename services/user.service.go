package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
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

// TODO: Implement
func (service *UserService) CreateUser(user *models.User) error {
	return nil
}

func (service *UserService) GetUsers() ([]models.User, error) {
	var users []models.User

	rows, err := service.db.Query(`
	SELECT id, username, roles FROM users;
	`)
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var roles string

		err := rows.Scan(&user.Id, &user.Username, &roles)
		if err != nil {
			// TODO: Better error handling
			fmt.Println(err)
			continue
		}

		user.Roles = utils.ParseRoles(roles)
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}

	return users, nil
}

func (service *UserService) UpdateUser(uid string) error {
	return nil
}

func (service *UserService) DeleteUser(uid string) error {
	return nil
}

func (service *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	var roles string

	row := service.db.QueryRow(`
	SELECT id, password_hash, roles FROM users
		WHERE username=$1
	`, username)

	err := row.Scan(&user.Id, &user.PasswordHash, &roles)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	user.Username = username
	user.Roles = utils.ParseRoles(roles)

	return &user, nil
}

func (service *UserService) GetUserById(id string) (*models.User, error) {
	var user models.User
	var roles string

	row := service.db.QueryRow(`
	SELECT username, password_hash, roles FROM users
		WHERE id=$1
	`, id)

	err := row.Scan(&user.Id, &user.PasswordHash, &roles)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	user.Id = id
	user.Roles = utils.ParseRoles(roles)

	return &user, nil
}
