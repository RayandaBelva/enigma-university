package repository

import (
	"database/sql"
	"time"

	"github.com/eulbyvan/enigma-university/model"
)

type UserRepository interface {
	// create
	Create(user model.User) error

	// read
	GetById(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)

	// update
	UpdateById(id string, updatedUser model.User) error

	// delete
	DeleteById(id string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) GetById(id string) (model.User, error) {
	// implementasikan cara get by id
	var user model.User
	query := `	SELECT 	id,
						first_name,
						last_name,
						email,
						username,
						role,
						photo,
						created_at,
						updated_at
				FROM	users
				WHERE	id = $1`

	err := u.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
		&user.Role,
		&user.Photo,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return []model.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Role, &user.Photo, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return []model.User{}, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (u *userRepository) Create(user model.User) error {
	query := `INSERT INTO users (first_name, last_name, email, username, role, photo, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id `

	err := u.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Username,
		user.Role,
		user.Photo,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UpdateById(id string, updatedUser model.User) error {

	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3, username = $4, role = $5, photo = $6, updated_at = $7 WHERE id = $8"
	_, err := u.db.Exec(query, updatedUser.FirstName, updatedUser.LastName, updatedUser.Email, updatedUser.Username, updatedUser.Role, updatedUser.Photo, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) DeleteById(id string) error {

	query := "DELETE FROM users WHERE id = $1"
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// constructor
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
