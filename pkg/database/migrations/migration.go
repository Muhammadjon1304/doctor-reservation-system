package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

const dbURL = "postgres://muhammadjonparpiyev:root@localhost:5432/doctor_reservations?sslmode=disable"

func RunMigrations() {
	m, err := migrate.New(
		"file://pkg/database/migrations", // Path to migration files
		dbURL,
	)
	if err != nil {
		log.Fatal("Failed to create migration instance:", err)
	}

	// Run Migrations Up
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Migrations applied successfully!")
}
