package service

import (
	"github.com/Feruz666/neuromaps-auth/models"
	"github.com/Feruz666/neuromaps-auth/pkg/repository"
)

// Authorization ...
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// Service ...
type Service struct {
	Authorization
}

// NewService ...
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
