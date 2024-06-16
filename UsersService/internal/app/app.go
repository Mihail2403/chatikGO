package app

import (
	"context"
	"log"
	"os"
	"time"
	"users-service/internal/repository"
	"users-service/internal/server"
	"users-service/internal/service"
	handler "users-service/internal/transport/http"
)

// func initEnv() {
// 	// loads values from .env into the system
// 	if err := godotenv.Load(); err != nil {
// 		log.Printf("No .env file found: %v", err)
// 	}
// }

func Run() {
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

	mongoCfg := repository.MongoConfig{
		Host: os.Getenv("MONGO_HOST"),
		Port: os.Getenv("MONGO_PORT"),
		DB:   os.Getenv("MONGO_DB"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	mongo, err := repository.NewMongoDB(ctx, mongoCfg)
	if err != nil {
		log.Fatalf("Error get mongo.Database: %v\n", err)
	}

	// dependency injection
	repo := repository.NewRepository(db, mongo)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	// init and run server
	srv := new(server.Server)
	err = srv.Run(os.Getenv("API_PORT"), handler.InitRoutes())
	if err != nil {
		log.Fatalln(err)
	}
}
