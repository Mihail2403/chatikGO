package repository

import (
	"auth/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

// constructor auth repository
func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

// method to insert  user data into database table 'users'
func (r *AuthRepo) SignUp(usr *entity.User) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (username, password_hash)"+
			"VALUES ($1, $2)",
		UsersTable,
	)
	_, err := r.db.Exec(query, usr.Username, usr.PasswordHash)
	return err
}

// method to get one user by id from database table 'users'
func (r *AuthRepo) GetUserById(id int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id, username, password_hash FROM %s WHERE id = $1 LIMIT 1",
		UsersTable,
	)
	err := r.db.Get(&user, query, id)
	return user, err
}

// method to get one user by password and username from the database table 'users'
func (r *AuthRepo) GetUserByUnameAndPasswordHash(username, password_hash string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id, username, password_hash FROM %s WHERE username = $1 AND password_hash = $2 LIMIT 1",
		UsersTable,
	)
	err := r.db.Get(&user, query, username, password_hash)
	return user, err
}
