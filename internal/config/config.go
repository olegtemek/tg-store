package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env                string `env:"ENV"`
	DatabaseUrl        string `env:"DATABASE_URL"`
	AccessTokenSecret  string `env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `env:"REFRESH_TOKEN_SECRET"`
	GRPC               GRPCConfig
}

type GRPCConfig struct {
	Port    string        `env:"GRPC_PORT"`
	Timeout time.Duration `env:"GRPC_TIMEOUT"`
}

func New() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		panic("Cannot read env")
	}

	return &cfg
}
