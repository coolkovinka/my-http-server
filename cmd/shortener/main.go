package main

import (
	"fmt"
	"my-http-server/config"
	"my-http-server/internal/app/handlers"
	"my-http-server/internal/pkg/server"
	"my-http-server/internal/pkg/storage"
)

func main() {
	// app config
	cfg := config.NewConfig()

	strConfig, _ := cfg.Sprint()
	fmt.Printf("[config]: %+v\n", strConfig)

	// storage
	repo := storage.NewStorage()

	// handlers
	serviceHandlers := handlers.NewHandlers(cfg, repo)

	// server
	newServer := server.NewServer(serviceHandlers, cfg)

	err := newServer.Start()
	if err != nil {
		panic(err)
	}

}
