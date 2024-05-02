package handlers

import "my-http-server/config"

type Storage interface {
	GetByURLPath(string) string
	SetByOriginalURL(string) string
}

type Handlers struct {
	cfg     *config.Config
	storage Storage
}

func NewHandlers(cfg *config.Config, storage Storage) *Handlers {
	return &Handlers{
		storage: storage,
		cfg:     cfg,
	}
}
