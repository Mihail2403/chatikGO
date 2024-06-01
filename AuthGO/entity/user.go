package entity

import (
	"auth/pkg/hash"
	"errors"
)

// Entity
type User struct {
	Id           int    `db:"id" json:"id"`
	Username     string `db:"username" json:"username"`
	PasswordHash string `db:"password_hash" json:"-"`
}

type UserFromJSON struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserFromJSON) ToEntity() (*User, error) {
	if u.Username == "" || u.Password == "" {
		return nil, errors.New("username or password  is empty")
	}
	password_hash := hash.GeneratePasswordHash(u.Password)
	return &User{
		Id:           u.Id,
		Username:     u.Username,
		PasswordHash: password_hash,
	}, nil
}
