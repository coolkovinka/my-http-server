package main

import (
	"log"

	"my-http-server/internal/app/handlers"
	"my-http-server/internal/config"
	"my-http-server/internal/pkg/server"
	"my-http-server/internal/pkg/storage"
)

func main() {
	// app config
	cfg := config.NewConfig()

	// storage
	repo := storage.NewStorage()

	// handlers
	serviceHandlers := handlers.NewHandlers(cfg, repo)

	// server
	newServer := server.NewServer(serviceHandlers, cfg)

	err := newServer.Start()
	if err != nil {
		log.Fatalf("fail staring server %v", err)
	}
}
