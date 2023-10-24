package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *Config) (*sql.DB, error) {
	// "postgres://username:password@localhost/database_name?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Name, cfg.SSL)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
