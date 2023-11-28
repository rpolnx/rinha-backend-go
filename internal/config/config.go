package configs

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/samber/do"
)

type EnvConfig struct {
	Port       int    `env:"PORT" envDefault:"8080"`
	DbHost     string `env:"DB_HOST"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbDbname   string `env:"DB_DBNAME"`
	DbPort     int    `env:"DB_PORT"`
	DbSslmode  string `env:"DB_SSLMODE" envDefault:"disable"`
	DbTimezone string `env:"DB_TIMEZONE" envDefault:"DB_TIMEZONE"`
}

func NewEnvConfig(injector *do.Injector) (*EnvConfig, error) {
	cfg := EnvConfig{}
	err := env.Parse(&cfg)
	if err != nil {
		err = fmt.Errorf("error serializing envs: %v", err)
	}

	return &cfg, err
}
