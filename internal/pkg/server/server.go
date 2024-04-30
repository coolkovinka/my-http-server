package server

import (
	"net/http"
)

type Handlers interface {
	SetShortURL(response http.ResponseWriter, request *http.Request)
	GetOriginalURL(response http.ResponseWriter, request *http.Request)
}

// Server represents the server to act.
type Server struct {
	handlers Handlers
}

// NewServer creates a new server .
func NewServer(handlers Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

// Start starts the server
func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc(`/`, s.handlers.SetShortURL)
	mux.HandleFunc(`/{id}`, s.handlers.GetOriginalURL)

	return http.ListenAndServe(`localhost:8080`, mux)
}
