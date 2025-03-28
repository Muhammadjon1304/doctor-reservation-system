package auth

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *User) error {
	// Hash password
	if err := user.HashPassword(); err != nil {
		return err
	}

	query := `
		INSERT INTO users (username, password, email, phone, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		user.Username,
		user.Password,
		user.Email,
		user.Phone,
		time.Now(),
	).Scan(&user.ID)

	return err
}

func (r *Repository) FindByUsername(username string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, username, password, email, phone, created_at
		FROM users 
		WHERE username = $1
	`

	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
