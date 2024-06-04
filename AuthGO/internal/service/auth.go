package service

import (
	"auth/entity"
	"auth/internal/repository"
	myjwt "auth/pkg/jwt"
	"fmt"
)

// Auth Interface Implementation
type AuthService struct {
	repo repository.Auth
}

// AuthService Constructor
func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repo,
	}
}

// Register a new user to the system
func (s *AuthService) SignUp(usr *entity.User) error {
	return s.repo.SignUp(usr)
}

// Login an existing user and generate JWT
func (s *AuthService) GetToken(usr *entity.User) (string, error) {
	u, err := s.repo.GetUserByUnameAndPasswordHash(usr.Email, usr.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %v", err)
	}

	token, err := myjwt.GenerateToken(u.Id)
	if err != nil {
		return "", fmt.Errorf("invalid token generating: %v", err)
	}

	return token, nil
}

// Get User by Token
func (s *AuthService) GetUserByToken(token string) (entity.User, error) {
	id, err := myjwt.ParseToken(token)
	if err != nil {
		return entity.User{}, fmt.Errorf("parse token failed: %v", err)
	}

	user, err := s.repo.GetUserById(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to get user by id: %v", err)
	}

	return user, nil
}
