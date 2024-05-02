package main

import (
	"fmt"

	"my-http-server/internal/app/handlers"
	"my-http-server/internal/pkg/config"
	"my-http-server/internal/pkg/server"
	"my-http-server/internal/pkg/storage"
)

func main() {

	cfg := config.NewConfig()
	strConfig, err := cfg.Sprint()
	if err != nil {
		fmt.Printf("print config error %s", err)
	}
	fmt.Printf("[config]: %+v\n", strConfig)

	// storage
	repo := storage.NewStorage()

	// handlers
	serviceHandlers := handlers.NewHandlers(repo)

	// server
	newServer := server.NewServer(serviceHandlers, cfg)
	err = newServer.Start()
	if err != nil {
		panic(err)
	}

}
