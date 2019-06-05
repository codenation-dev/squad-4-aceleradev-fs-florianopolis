package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "uati"
	SSLMODE     = "disable"
)

func Connect() *sql.DB {
	URL := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, DB_NAME, SSLMODE,
	)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		// return fmt.Errorf("%v", err)
	}
	fmt.Println("Database connected!!!")
	// return nil
}
