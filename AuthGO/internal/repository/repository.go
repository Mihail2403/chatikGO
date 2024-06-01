package repository

import (
	"auth/entity"

	"github.com/jmoiron/sqlx"
)

// Interface of Auth Repository
type Auth interface {
	SignUp(usr *entity.User) error
	GetUserById(id int) (entity.User, error)
	GetUserByUnameAndPasswordHash(username, password_hash string) (entity.User, error)
}

// struct  for implementing the Repository interfaces
type Repository struct {
	Auth
}

// Constructor of Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
	}
}
