package app

import (
	"chat/internal/server"
	handler "chat/internal/transport/http"
	"os"
)

//	func initEnv() {
//		// loads values from .env into the system
//		if err := godotenv.Load(); err != nil {
//			log.Printf("No .env file found: %v", err)
//		}
//	}

func Run() {
	handler := handler.NewHandler().InitRoutes()
	srv := new(server.Server)
	srv.Run(os.Getenv("API_PORT"), handler)
}
