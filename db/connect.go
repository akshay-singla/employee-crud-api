package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// NewConnection creates connection to MySQL or MariaDB database.
func NewConnection(cfg *Config) (*sql.DB, error) {
	log.Println(cfg.DSN())
	conn, err := sql.Open(cfg.Driver, cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("open sql connection: %w", err)
	}
	return conn, nil
}

// Configure prepares database connection for use in app.
func Configure(conn *sql.DB, config *Config) error {
	stmt := fmt.Sprintf(`SET SESSION TRANSACTION ISOLATION LEVEL %s`, sql.LevelReadCommitted)
	if _, err := conn.Exec(stmt); err != nil {
		return fmt.Errorf("set tx iso level: %w", err)
	}

	conn.SetMaxOpenConns(int(config.MaxConns))
	conn.SetMaxIdleConns(int(config.IdleConns))
	conn.SetConnMaxIdleTime(config.IdleConnLifetime)

	return nil
}

// Verify ensures the connection is available for use.
func Verify(conn *sql.DB) error {
	if err := conn.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}

func CreateTables(conn *sql.DB) {
	// Create employees table if not exists
	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS employees (
	id SERIAL PRIMARY KEY,
	name TEXT,
	position TEXT,
	salary NUMERIC
)`)
	if err != nil {
		log.Fatal(err)
	}
}
