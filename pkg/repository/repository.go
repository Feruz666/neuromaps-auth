package repository

import (
	"github.com/Feruz666/neuromaps-auth/models"
	"github.com/jmoiron/sqlx"
)

// Authorization ...
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

// Repository ...
type Repository struct {
	Authorization
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
