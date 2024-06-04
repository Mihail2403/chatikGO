package server

import (
	"fmt"
	"net/http"
	"os"
)

var (
	AUTH_URL = fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_HOST"), os.Getenv("AUTH_PORT"))
)

// base Server struct
type Server struct {
	httpServer *http.Server
}

// func for  creating a new instance of the http server with port and handler
func (s *Server) Run(port string, handler http.Handler) error {
	// create  the HTTP server
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// run the  HTTP server
	return s.httpServer.ListenAndServe()
}
