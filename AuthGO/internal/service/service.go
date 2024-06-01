package service

import (
	"auth/entity"
	"auth/internal/repository"
)

// Auth Service is the interface for handling authentication related operations
type Auth interface {
	SignUp(usr *entity.User) error
	GetToken(usr *entity.User) (string, error)
	GetUserByToken(token string) (entity.User, error)
}

// Service  implements the Services interfaces
type Service struct {
	Auth
}

// Service Constructor
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
