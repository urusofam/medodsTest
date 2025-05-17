package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	DatabaseConfig struct {
		Host     string `env:"DB_HOST" env-default:"postgres"`
		Port     int    `env:"DB_PORT" env-default:"5432"`
		User     string `env:"DB_USER" env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"admin123"`
		Database string `env:"DB_DATABASE" env-default:"authDB"`
	}

	ServerConfig struct {
		Address      string        `env:"SERVER_ADDRESS" env-default:"localhost:8080"`
		ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" env-default:"5s"`
		WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" env-default:"10s"`
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading .env: %v", err)
	}

	return &cfg, nil
}
