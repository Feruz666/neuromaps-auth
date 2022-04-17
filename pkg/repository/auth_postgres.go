package repository

import (
	"fmt"

	"github.com/Feruz666/neuromaps-auth/models"
	"github.com/jmoiron/sqlx"
)

// AuthPostgres ...
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres ...
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser ...
func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", userTable)
	err := r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetUser ...
func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
