package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AkifhanIlgaz/key-aero-api/errors"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/AkifhanIlgaz/key-aero-api/utils"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type UserService struct {
	psql squirrel.StatementBuilderType
	ctx  context.Context
}

func NewUserService(ctx context.Context, db *sql.DB) *UserService {
	return &UserService{
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
		ctx:  ctx,
	}
}

func (service *UserService) CreateUser(input models.User) error {
	passwordHash, err := utils.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	_, err = service.psql.Insert("users").Columns("username", "password_hash", "roles", "email", "phone", "department").Values(input.Username, passwordHash, utils.GenerateRolesString(input.Roles), input.Email, input.Phone, input.Department).Exec()
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

	rows, err := service.psql.Select("id", "username", "roles", "email", "phone", "department").From("users").Query()
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		var roles string

		err := rows.Scan(&user.Id, &user.Username, &roles, &user.Email, &user.Phone, &user.Department)
		if err != nil {
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
func (service *UserService) UpdateUser(id string, updates map[string]any) error {
	_, err := service.psql.Update("users").Where(squirrel.Eq{"id": id}).SetMap(updates).Exec()
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}

func (service *UserService) SearchUser(search models.SearchUserInput) ([]models.User, error) {
	var users []models.User

	/*
		rows, err := service.db.Query(`
			SELECT id, username, roles, email, phone, department FROM users
			WHERE username LIKE CONCAT('%', $1::TEXT, '%')
			 AND roles LIKE CONCAT('%', $2::TEXT, '%')
			 AND email LIKE CONCAT('%', $3::TEXT, '%')
			 AND phone LIKE CONCAT('%', $4::TEXT, '%')
			 AND department LIKE CONCAT('%', $5::TEXT, '%');
		`, search.Username, search.Roles, search.Email, search.Phone, search.Department)
		if err != nil {
			return nil, fmt.Errorf("search user: %w", err)
		}
		defer rows.Close()
	*/

	rows, err := service.psql.Select("id", "username", "roles", "email", "phone", "department").From("users").Where(squirrel.And{
		squirrel.Like{"username": "%" + search.Username + "%"},
		squirrel.Like{"roles": "%" + search.Roles + "%"},
		squirrel.Like{"email": "%" + search.Email + "%"},
		squirrel.Like{"phone": "%" + search.Phone + "%"},
		squirrel.Like{"department": "%" + search.Department + "%"},
	}).Query()
	if err != nil {
		return nil, fmt.Errorf("search user: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		var roles string

		err := rows.Scan(&user.Id, &user.Username, &roles, &user.Email, &user.Phone, &user.Department)
		if err != nil {
			fmt.Println(err)
			continue
		}

		user.Roles = utils.ParseRoles(roles)
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("search user: %w", err)
	}

	return users, nil
}

func (service *UserService) DeleteUser(uid string) error {
	_, err := service.psql.Delete("users").Where(squirrel.Eq{"id": uid}).Exec()

	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

func (service *UserService) GetUserByUsername(username string) (*models.User, error) {
	row := service.psql.Select("id", "password_hash", "roles", "email", "phone", "department").From("users").Where(squirrel.Eq{
		"username": username,
	}).QueryRow()

	var user models.User
	var roles string

	err := row.Scan(&user.Id, &user.PasswordHash, &roles, &user.Email, &user.Phone, &user.Department)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	user.Username = username
	user.Roles = utils.ParseRoles(roles)

	return &user, nil
}

func (service *UserService) GetUserById(id string) (*models.User, error) {
	row := service.psql.Select("username", "password_hash", "roles", "email", "phone", "department").From("users").Where(squirrel.Eq{
		"id": id,
	}).QueryRow()

	var user models.User
	var roles string

	err := row.Scan(&user.Username, &user.PasswordHash, &roles, &user.Email, &user.Phone, &user.Department)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	user.Id = id
	user.Roles = utils.ParseRoles(roles)

	return &user, nil
}
