package repository

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
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
	db, err := sqlx.Open("pgx",
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
