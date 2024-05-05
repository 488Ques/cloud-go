package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Server ConfigServer
	DB     ConfigDB
}

type ConfigServer struct {
	Port         int           `env:"SERVER_PORT"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE"`
	Debug        bool          `env:"SERVER_DEBUG"`
}

type ConfigDB struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	DBName   string `env:"DB_NAME"`
	Debug    bool   `env:"DB_DEBUG"`
}

func New() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %e", err)
	}

	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Unable to parse environment variable: %e", err)
	}

	return cfg
}
