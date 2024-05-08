package config

import (
	"os"
	"strconv"
	"time"

	"github.com/akshay-singla/employee-crud-api/db"
)

// GetEnvConfig retrieves environment variables and populates the Config struct
func GetEnvConfig() *db.Config {
	cfg := &db.Config{
		User:     getEnv("DB_USER", "company-test"),
		Password: getEnv("DB_PASSWORD", "company-test"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Name:     getEnv("DB_NAME", "company-test"),
		Driver:   getEnv("DB_DRIVER", "postgres"),
	}

	maxConns, err := strconv.ParseUint(getEnv("MAX_CONNS", "10"), 10, 32)
	if err == nil {
		cfg.MaxConns = uint(maxConns)
	}

	idleConns, err := strconv.ParseUint(getEnv("IDLE_CONNS", "10"), 10, 32)
	if err == nil {
		cfg.IdleConns = uint(idleConns)
	}

	idleConnLifetime, err := time.ParseDuration(getEnv("IDLE_CONN_LIFETIME", "5m"))
	if err == nil {
		cfg.IdleConnLifetime = idleConnLifetime
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
