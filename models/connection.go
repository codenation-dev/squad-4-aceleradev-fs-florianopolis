package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // postgres
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "uati"
	SSLMODE     = "disable"
	HOST        = "172.17.0.2"
	PORT        = "5432"
)

// Connect to the db
func Connect() *sql.DB {
	URL := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, HOST, PORT, DB_NAME, SSLMODE,
	)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// TestConnection with the db
func TestConnection() {
	db := Connect()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println("Database connected!!!")
}
