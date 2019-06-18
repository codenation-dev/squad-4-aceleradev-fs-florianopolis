package main

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres
)

type repo struct {
	db *sql.DB
}

// ReadCustomers read customers from the DB
// If id = 0, it will return all customers
func (r *repo) Read(id int) (*Customer, error) {
	var c Customer

	query := "SELECT * FROM customers WHERE id=$1"
	err := r.db.QueryRow(query, id).Scan(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
