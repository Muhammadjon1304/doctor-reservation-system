package database

import (
	"database/sql"
	"fmt"

	cfg "doctor-reservation-system/config"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg cfg.DatabaseConfig) (*sql.DB, error) {
	// Construct connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	// Open connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Verify connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}
