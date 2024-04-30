package config

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"HTTP_SERVER_PORT"`
	Host string `envconfig:"HTTP_SERVER_HOST"`
}

// NewConfig returns app config
func NewConfig() *Config {
	err := godotenv.Load("infra/config.env")
	if err != nil {
		panic(err)
	}

	cfg := new(Config)

	err = envconfig.Process("", cfg)
	if err != nil {
		fmt.Printf("envconfig err: %v\n", err.Error())
	}

	return cfg
}

func (c *Config) Sprint() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
