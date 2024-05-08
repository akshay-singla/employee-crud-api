package db

import (
	"fmt"
	"time"
)

type Config struct {
	User     string `default:"company-test"`
	Password string `default:"company-test"`
	Host     string `default:"localhost"`
	Port     string `default:"5432"`
	Name     string `default:"company-test"`
	Driver   string `default:"postgres"`

	MaxConns         uint          `default:"10"`
	IdleConns        uint          `default:"10"`
	IdleConnLifetime time.Duration `default:"5m"`
}

func (cfg *Config) DSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Driver, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
