package models

import (
	// "codenation/squad-4-aceleradev-fs-florianopolis/models"

	"fmt"
)

// Customer of the bank
type Customer struct {
	CID         int32   `json:"cid"`
	Name        string  `json:"name"`
	Wage        float32 `json:"wage"`
	IsPublic    int8    `json:"is_public"`
	SentWarning string  `json:"sent_warning"`
}

// NewCustomer insert a new customer on the db
func NewCustomer(c Customer) (bool, error) {
	db := Connect()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	sql := `INSERT INTO customers (name, wage, is_public, sent_warning)
			VALUES ($1, $2, $3, $4)
			RETURNING CID`
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		err = stmt.QueryRow(
			c.Name, c.Wage, c.IsPublic, c.SentWarning,
		).Scan(&c.CID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit()
}

// GetCustomers returns a list of all bank's customers
func GetCustomers() ([]Customer, error) {
	db := Connect()
	defer db.Close()
	sql := "SELECT * FROM customers"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.CID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// GetCustomersPublicFuncs returns the bank customers that are public employees
func GetCustomersPublicFuncs() ([]Customer, error) {
	db := Connect()
	defer db.Close()
	sql := "SELECT * FROM customers WHERE is_public = $1"
	rows, err := db.Query(sql, 1)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.CID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// GetVIPCustomers returns the bank customers that earns a good wage
func GetVIPCustomers(goodWage float32) ([]Customer, error) {
	db := Connect()
	defer db.Close()
	sql := "SELECT * FROM customers WHERE wage > $1"
	rows, err := db.Query(sql, goodWage)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.CID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
