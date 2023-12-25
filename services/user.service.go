package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

func (service *UserService) CreateUser(input models.User) error {
	passwordHash, err := utils.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	_, err = service.db.Exec(`
		INSERT INTO users (username, password_hash, roles, email, phone, department)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		);
	`, input.Username, passwordHash, input.Roles, input.Email, input.Phone, input.Department)

	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == pgerrcode.UniqueViolation {
				return errors.ErrUsernameTaken
			}
		}
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (service *UserService) GetUsers() ([]models.User, error) {
	var users []models.User

	rows, err := service.db.Query(`
	SELECT id, username, roles, email, phone, department FROM users;
	`)
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var roles string

		err := rows.Scan(&user.Id, &user.Username, &roles, &user.Email, &user.Phone, &user.Department)
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

// Tüm kullanıcı bilgilerini gönder
func (service *UserService) UpdateUser(updatedUser *models.UpdateInput) error {
	_, err := service.db.Exec(`
		UPDATE users
		SET 
		username = $2,
		roles = $3,
		email = $4,
		phone = $5,
		department = $6
		WHERE id = $1;
	`, updatedUser.Id, updatedUser.Username, utils.GenerateRolesString(updatedUser.Roles), updatedUser.Email, updatedUser.Phone, updatedUser.Department)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}

// TODO: Implement this function
func (service *UserService) SearchUser(search []gin.Param) ([]models.User, error) {

	return nil, nil
}

func (service *UserService) DeleteUser(uid string) error {
	_, err := service.db.Exec(`
		DELETE FROM users
		WHERE id = $1;
	`, uid)

	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

func (service *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	var roles string

	row := service.db.QueryRow(`
	SELECT id, password_hash, roles, email, phone, department FROM users
		WHERE username=$1
	`, username)

	err := row.Scan(&user.Id, &user.PasswordHash, &roles, &user.Email, &user.Phone, &user.Department)
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
	SELECT username, password_hash, roles, email, phone, department FROM users
		WHERE id=$1
	`, id)

	err := row.Scan(&user.Id, &user.PasswordHash, &roles, &user.Email, &user.Phone, &user.Department)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	user.Id = id
	user.Roles = utils.ParseRoles(roles)

	return &user, nil
}
