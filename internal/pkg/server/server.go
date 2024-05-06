package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"my-http-server/internal/config"
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

// Start starts the server.
func (s *Server) Start() error {
	router := chi.NewRouter()

	router.Get(`/{id}`, s.handlers.GetOriginalURL)
	router.Post(`/`, s.handlers.SetShortURL)

	log.Printf("Running server on %s", s.cfg.ServerAddress)

	err := http.ListenAndServe(s.cfg.ServerAddress, router)
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return errors.New(err.Error())
		}
	}

	return nil
}
