package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	_ "github.com/lib/pq" // postgres
)

const (
	// DB data
	DRIVER_NAME = "postgres"
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	// DB_NAME     = "uati"
	SSLMODE = "disable"
	HOST    = "172.17.0.2"
	PORT    = "5432"
)

// Storage stores data ia a postgresql db
type Storage struct {
	db *sql.DB
}

// Connect implements the connection to the db
func Connect(dbName string) (*sql.DB, error) {
	// var err error
	connString := fmt.Sprintf(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, HOST, PORT, dbName, SSLMODE,
	))

	return sql.Open("postgres", connString)
	// if err != nil {
	// 	panic(err)
	// }
	// return db
}

// NewStorage creates a new instance of Storage
func NewStorage(dbName string) *Storage {
	db, err := Connect(dbName)
	if err != nil {
		log.Fatal(err)
	}
	s := new(Storage)
	s.db = db
	return s
}

// ReadUser reads a customer from the DB, given the email
func (s *Storage) ReadUser(email string) (entity.User, error) {
	user := entity.User{}
	user.Email = email
	query := "SELECT password FROM users WHERE email=$1"
	err := s.db.QueryRow(query, email).Scan(&user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, entity.ErrUserNotFound
		}
		return entity.User{}, err
	}
	return user, err
}

// UpdateUser replace some data mantaining the same id
func (s *Storage) UpdateUser(u entity.User) error {
	res, err := s.db.Exec(`UPDATE users
						SET password=$1
						WHERE email=$2`, &u.Password, &u.Email)
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		err = entity.ErrUserNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
