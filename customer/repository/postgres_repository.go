package repository

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/customer"
	"codenation/squad-4-aceleradev-fs-florianopolis/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

type repo struct {
	db *sql.DB
}

// NewRepo creates a new instance of repo
func NewRepo(db *sql.DB) customer.Repository {
	return &repo{db}
}

// Create inserts a new customer on the DB
func (r *repo) Create(c *models.Customer) (*models.Customer, error) {
	err := r.db.QueryRow(`INSERT INTO customers (name, wage, is_public, sent_warning)
							VALUES ($1, $2, $3, $4)
							RETURNING id`, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning).Scan(&c.ID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Read read customers from the DB
func (r *repo) Read(id int) (*models.Customer, error) {
	var c models.Customer
	query := "SELECT * FROM customers WHERE id=$1"
	err := r.db.QueryRow(query, id).Scan(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// ReadAll return all customers from the DB
func (r *repo) ReadAll() ([]*models.Customer, error) {
	customers := []*models.Customer{}
	rows, err := r.db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := models.Customer{}
		err := rows.Scan(&c)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}

// ReadByName returns a slice of customers with the given pattern in the name column
func (r *repo) ReadByName(pattern string) ([]*models.Customer, error) {
	customers := []*models.Customer{}
	rows, err := r.db.Query(`SELECT * FROM customers WHERE name LIKE "%$1%"`, pattern)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := models.Customer{}
		err := rows.Scan(&c)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}

// Update replace some data mantaining the same id
func (r *repo) Update(c *models.Customer) error {
	_, err := r.db.Exec(`UPDATE customers
						SET name=$1, wage=$2, is_public=$3, sent_warning=$4
						WHERE id=$5`, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning, &c.ID)
	if err != nil {
		return fmt.Errorf("could not delete the customer: %v", err)
	}
	return nil
}

func (r *repo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM customers WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("could not delete customer: %v", err)
	}
	return nil
}
