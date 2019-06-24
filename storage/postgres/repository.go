package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

const (
	// DB data
	DRIVER_NAME = "postgres"
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "uati"
	SSLMODE     = "disable"
	HOST        = "172.17.0.2"
	PORT        = "5432"
)

// Storage stores data ia a postgresql db
type Storage struct {
	db *sql.DB
}

func Connect() *sql.DB {
	var err error
	connString := fmt.Sprintf(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, HOST, PORT, DB_NAME, SSLMODE,
	))

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	return db
}

// NewStorage creates a new instance of Storage
func NewStorage(db *sql.DB) (*Storage, error) {
	s := new(Storage)
	s.db = db
	return s, nil
}
