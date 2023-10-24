package db

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	// "postgres://username:password@localhost/database_name?sslmode=disable"
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE")))
	if err != nil {
		return nil, err
	}
	return db, nil
}
