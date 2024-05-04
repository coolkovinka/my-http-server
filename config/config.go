package config

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	defaultServerAddress = "localhost:8080"
	defaultBaseUrl       = "http://localhost:8080"
)

type Config struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS" default:""`
	BaseUrl       string `envconfig:"BASE_URL" default:""`
}

// NewConfig returns app config
func NewConfig() *Config {
	cfg := new(Config)

	err := envconfig.Process("", cfg)
	if err != nil {
		fmt.Printf("error processing env config %v", err)
	}

	if cfg.ServerAddress == "" || cfg.BaseUrl == "" {
		cfg.parseFlags()
	}

	if cfg.ServerAddress == "" || cfg.BaseUrl == "" {
		fmt.Print("default env config is set \n")
		cfg.setDefault()
	}

	return cfg
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.ServerAddress, "a", "", "address and port to run server")
	flag.StringVar(&c.BaseUrl, "b", "", "base URL with short URL")

	flag.Parse()
}

func (c *Config) Sprint() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Config) setDefault() {
	c.ServerAddress = defaultServerAddress
	c.BaseUrl = defaultBaseUrl
}
