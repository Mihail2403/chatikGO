package app

import (
	"auth/internal/repository"
	"auth/internal/server"
	"auth/internal/service"
	handler "auth/internal/transport/http"
	"log"
	"os"
)

// func initEnv() {
// 	// loads values from .env into the system
// 	if err := godotenv.Load(); err != nil {
// 		log.Printf("No .env file found: %v", err)
// 	}
// }

func Run() {
	// init postgres
	cfg := repository.Config{
		Port:     os.Getenv("PG_PORT"),
		Password: os.Getenv("PG_PASSWORD"),
		Host:     os.Getenv("PG_HOST"),
		DBName:   os.Getenv("PG_DATABASE"),
		Username: os.Getenv("PG_USERNAME"),
	}
	db, err := repository.NewPostgresDB(&cfg)
	if err != nil {
		log.Fatalf("Error get sql.DB: %v\n", err)
	}
	defer db.Close()

	// dependency injection
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	// init and run server
	srv := new(server.Server)
	err = srv.Run(os.Getenv("API_PORT"), handler.InitRoutes())
	if err != nil {
		log.Fatalln(err)
	}
}
