package config

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Server configServer
	DB     configDB
}

type configServer struct {
	Host         string        `env:"SERVER_HOST"`
	Port         int           `env:"SERVER_PORT"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE"`
	Debug        bool          `env:"SERVER_DEBUG"`
}

type configDB struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	DBName   string `env:"DB_NAME"`
	Debug    bool   `env:"DB_DEBUG"`
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %e", err)
	}

	cfg := &Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Unable to parse environment variable: %e", err)
	}

	return cfg
}

func (c *Config) NewConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName)
}
