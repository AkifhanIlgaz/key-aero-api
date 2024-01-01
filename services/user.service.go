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

	_, err = service.psql.Insert("users").Columns("username", "password_hash", "roles", "email", "phone", "department", "name").Values(input.Username, passwordHash, input.Roles, input.Email, input.Phone, input.Department, input.Name).Exec()
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

	rows, err := service.psql.Select("id", "username", "roles", "email", "phone", "department", "name").From("users").Query()
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Username, &user.Roles, &user.Email, &user.Phone, &user.Department, &user.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}

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

func (service *UserService) ChangePassword(id string, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("change password: %w", err)
	}

	res, err := service.psql.Update("users").Where(squirrel.Eq{"id": id}).Set("password_hash", hashedPassword).Exec()
	if err != nil {
		return fmt.Errorf("change password: %w", err)
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("change password: %w", err)
	}

	return nil
}

// TODO: Her rolü ayrı ara || Sırası farklı olabilir
// TODO: Name ekle
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

	roles := utils.ParseRoles(search.Roles)
	where := squirrel.And{
		squirrel.Like{"username": "%" + search.Username + "%"},
		squirrel.Like{"email": "%" + search.Email + "%"},
		squirrel.Like{"phone": "%" + search.Phone + "%"},
		squirrel.Like{"department": "%" + search.Department + "%"},
	}
	for _, role := range roles {
		where = append(where, squirrel.Like{"roles": "%" + role + "%"})
	}

	rows, err := service.psql.Select("id", "username", "roles", "email", "phone", "department").From("users").Where(where).Query()
	if err != nil {
		return nil, fmt.Errorf("search user: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Username, &user.Roles, &user.Email, &user.Phone, &user.Department)
		if err != nil {
			fmt.Println(err)
			continue
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("search user: %w", err)
	}

	return users, nil
}

func (service *UserService) DeleteUser(uids []string) error {
	_, err := service.psql.Delete("users").Where(squirrel.Eq{"id": uids}).Exec()
	// squirrel.Eq{"id": uid}
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

func (service *UserService) GetUserByUsername(username string) (*models.User, error) {
	row := service.psql.Select("id", "password_hash", "roles", "email", "phone", "department", "name").From("users").Where(squirrel.Eq{
		"username": username,
	}).QueryRow()

	var user models.User

	err := row.Scan(&user.Id, &user.PasswordHash, &user.Roles, &user.Email, &user.Phone, &user.Department, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	user.Username = username

	return &user, nil
}

func (service *UserService) GetUserById(id string) (*models.User, error) {
	row := service.psql.Select("username", "password_hash", "roles", "email", "phone", "department", "name").From("users").Where(squirrel.Eq{
		"id": id,
	}).QueryRow()

	var user models.User

	err := row.Scan(&user.Username, &user.PasswordHash, &user.Roles, &user.Email, &user.Phone, &user.Department, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	user.Id = id

	return &user, nil
}
