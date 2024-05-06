package config

import (
	"flag"
	"log"

	"github.com/kelseyhightower/envconfig"
)

const (
	defaultServerAddress = "localhost:8080"
	defaultBaseURL       = "http://localhost:8080"
)

type Config struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS" default:""`
	BaseURL       string `envconfig:"BASE_URL" default:""`
}

// NewConfig returns app config.
func NewConfig() *Config {
	cfg := new(Config)

	err := envconfig.Process("", cfg)
	if err != nil {
		log.Printf("error processing env config %v", err)
	}

	cfg.parseFlag()

	if cfg.ServerAddress == "" {
		cfg.setDefaultServerAddress()
	}

	if cfg.BaseURL == "" {
		cfg.setDefaultBaseURL()
	}

	return cfg
}

func (c *Config) parseFlag() {
	var serverAddress string
	var baseURL string

	flag.StringVar(&serverAddress, "a", "localhost:8080", "host and port to run server")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base URL with short URL")

	flag.Parse()

	if len(c.ServerAddress) == 0 {
		c.ServerAddress = serverAddress
	}
	if len(c.BaseURL) == 0 {
		c.BaseURL = baseURL
	}
}

func (c *Config) setDefaultServerAddress() {
	log.Print("default server address is set")

	c.ServerAddress = defaultServerAddress
}

func (c *Config) setDefaultBaseURL() {
	log.Print("default base URL is set")

	c.BaseURL = defaultBaseURL
}
