package memory

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
	"database/sql"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq" // postgres
)

// const (
// 	// DB data
// 	DRIVER_NAME = "postgres"
// 	DB_USER     = "postgres"
// 	DB_PASSWORD = "12345"
// 	DB_NAME     = "uati"
// 	SSLMODE     = "disable"
// 	HOST        = "172.17.0.2"
// 	PORT        = "5432"
// )

// Storage stores data ia a postgresql db
type Storage struct {
	db *sql.DB
}

// NewStorage creates a new instance of Storage
func NewStorage() (*Storage, error) {
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("could not create mock db: %v", err)
	}
	s := new(Storage)
	s.db = db
	return s, nil
}

// func NewStorage() (*Storage, error) {
// 	var err error
// 	s := new(Storage)
// 	connString := fmt.Sprintf(fmt.Sprintf(
// 		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
// 		DB_USER, DB_PASSWORD, HOST, PORT, DB_NAME, SSLMODE,
// 	))

// 	db, err := sql.Open("postgres", connString)
// 	s.db = db
// 	if err != nil {
// 		return nil, fmt.Errorf("could not connect to DB: %v", err)
// 	}
// 	return s, nil
// }

// DeleteCustomerByID delets a customer from the db
func (s *Storage) DeleteCustomerByID(id int) error {
	_, err := s.db.Exec(`DELETE FROM customers WHERE id=$1`, id)
	return err
}

// AddCustomer inserts a new customer on the DB
func (s *Storage) AddCustomer(c entity.Customer) error {
	_, err := s.db.Exec(`INSERT INTO customers (name, wage, is_public, sent_warning)
						VALUES ($1, $2, $3, $4)`,
		&c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
	return err
}

// GetCustomerByID read a customer from the DB, given the id
func (s *Storage) GetCustomerByID(id int) (entity.Customer, error) {
	c := entity.Customer{}
	query := "SELECT * FROM customers WHERE id=$1"
	err := s.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
	if err != nil {
		return entity.Customer{}, err
	}
	return c, err
}

// GetAllCustomers return all customers from the DB
func (s *Storage) GetAllCustomers() ([]entity.Customer, error) {
	customers := []entity.Customer{}
	rows, err := s.db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := entity.Customer{}
		err := rows.Scan(&c.ID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// GetCustomerByName returns a slice of customers with the given pattern in the name column
func (s *Storage) GetCustomerByName(pattern string) ([]entity.Customer, error) {
	customers := []entity.Customer{}
	query := fmt.Sprintf("SELECT * FROM customers WHERE name LIKE '%%%s%%'", pattern)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := entity.Customer{}
		err := rows.Scan(&c.ID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// UpdateCustomer replace some data mantaining the same id
func (s *Storage) UpdateCustomer(c entity.Customer) error {
	_, err := s.db.Exec(`UPDATE customers
						SET name=$1, wage=$2, is_public=$3, sent_warning=$4
						WHERE id=$5`, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning, &c.ID)
	if err != nil {
		return fmt.Errorf("could not update the customer: %v", err)
	}
	return nil
}
