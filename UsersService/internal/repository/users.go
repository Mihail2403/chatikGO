package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Create(user *UsersPostgresStruct) error {
	query := fmt.Sprintf("INSERT INTO %s (name, user_id, img_id) VALUES ($1, $2, $3)", UsersTable)
	_, err := r.db.Exec(query, user.Name, user.UserID, user.ImgID)
	return err
}
func (r *UsersRepo) GetAll() ([]UsersPostgresStruct, error) {
	var users []UsersPostgresStruct
	query := fmt.Sprintf("SELECT u.id, u.name, u.user_id, u.img_id FROM %s u", UsersTable)
	err := r.db.Select(&users, query)
	return users, err
}
func (r *UsersRepo) GetByID(id int) (UsersPostgresStruct, error) {
	var user UsersPostgresStruct
	query := fmt.Sprintf("SELECT u.id, u.name, u.user_id, u.img_id FROM %s u WHERE u.id=$1", UsersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *UsersRepo) Update(id int, user *UsersPostgresStruct) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, img_id=$2 WHERE id=$3", UsersTable)
	_, err := r.db.Exec(query, user.Name, user.ImgID, id)
	return err
}

func (r *UsersRepo) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s u WHERE u.id=$1", UsersTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UsersRepo) GetByIDArray(ids []int) ([]UsersPostgresStruct, error) {
	var users []UsersPostgresStruct
	query := fmt.Sprintf("SELECT u.id, u.name, u.user_id, u.img_id FROM %s u WHERE u.id = ANY($1)", UsersTable)
	err := r.db.Select(&users, query, ids)
	return users, err
}
