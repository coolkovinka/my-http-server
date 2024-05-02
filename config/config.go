package config

import (
	"encoding/json"
	"flag"
)

type Config struct {
	ServerAddress string
	ServerBaseURL string
}

func (c *Config) Sprint() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// NewConfig returns app config
func NewConfig() *Config {
	cfg := Config{}

	flag.StringVar(&cfg.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.ServerBaseURL, "b", "http://localhost:8080", "base URL with short URL")

	flag.Parse()

	return &cfg
}
