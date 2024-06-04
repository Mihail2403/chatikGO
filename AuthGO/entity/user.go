package entity

import (
	"auth/pkg/hash"
	"errors"
)

// Entity
type User struct {
	Id           int    `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
}

type UserFromJSON struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserFromJSON) ToEntity() (*User, error) {
	if u.Email == "" || u.Password == "" {
		return nil, errors.New("username or password  is empty")
	}
	password_hash := hash.GeneratePasswordHash(u.Password)
	return &User{
		Id:           u.Id,
		Email:        u.Email,
		PasswordHash: password_hash,
	}, nil
}
