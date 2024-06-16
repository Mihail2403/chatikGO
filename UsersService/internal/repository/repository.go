package repository

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersPostgresStruct struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	UserID int    `db:"user_id"`
	ImgID  string `db:"img_id"`
}

type Users interface {
	GetAll() ([]UsersPostgresStruct, error)
	GetByID(id int) (UsersPostgresStruct, error)
	Create(user *UsersPostgresStruct) error
	Update(id int, user *UsersPostgresStruct) error
	Delete(id int) error
	GetByIDArray(ids []int) ([]UsersPostgresStruct, error)
}
type Img interface {
	GetById(id string) (string, error)
	GetByIDArray(idArr []string) ([]string, error)
	Create(img string) (string, error)
	Update(id, img string) error
	Delete(id string) error
}

type Repository struct {
	Users
	Img
}

func NewRepository(db *sqlx.DB, mongoDB *mongo.Database) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
		Img:   NewImgRepo(mongoDB),
	}
}
