package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Name     string `env:"DB_NAME" default:"db_1"`
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     string `env:"DB_PORT" default:"5432"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type Origin struct {
	AllowOrigin string `env:"ALLOW_ORIGIN"`
}

type ServerConfig struct {
	ServerHost string `env:"SERVER_HOST" default:"localhost"`
	ServerPort string `env:"SERVER_PORT" default:"5000"`
	DB         Database
	Origin     Origin
	PrivateKey string `env:"PRIVATE_KEY"`
	PublicKey  string `env:"PUBLIC_KEY"`
}

var Config ServerConfig

func init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}
}

func loadConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn().Msg("Cannot find .env fils. OS Environtments will be used")
	}

	err := env.Parse(&Config)
	return err
}
