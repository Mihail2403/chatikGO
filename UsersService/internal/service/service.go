package service

import (
	"users-service/entity"
	"users-service/internal/repository"
)

type Users interface {
	GetAll() ([]entity.User, error)
	GetByID(id int) (*entity.User, error)
	Create(user *entity.User) error
	Update(id int, user *entity.User) error
	Delete(id int) error
	GetByIDArray(ids []int) ([]entity.User, error)
}

type Service struct {
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repo),
	}
}
