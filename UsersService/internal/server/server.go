package server

import (
	"net/http"
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
