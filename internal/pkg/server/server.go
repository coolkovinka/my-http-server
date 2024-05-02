package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"my-http-server/internal/pkg/config"
)

type Handlers interface {
	SetShortURL(response http.ResponseWriter, request *http.Request)
	GetOriginalURL(response http.ResponseWriter, request *http.Request)
}

// Server represents the server to act.
type Server struct {
	handlers Handlers
	cfg      *config.Config
}

// NewServer creates a new server .
func NewServer(handlers Handlers, cfg *config.Config) *Server {
	return &Server{
		handlers: handlers,
		cfg:      cfg,
	}
}

// Start starts the server
func (s *Server) Start() error {
	router := chi.NewRouter()

	router.Get(`/{id}`, s.handlers.GetOriginalURL)
	router.Post(`/`, s.handlers.SetShortURL)

	return http.ListenAndServe(s.cfg.Host+`:`+s.cfg.Port, router)
}
