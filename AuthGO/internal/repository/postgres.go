package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	UsersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	// creating connection to postgres
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password,
		),
	)
	if err != nil {
		return nil, err
	}

	// try to connect to db, its give the error if connection refused
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
