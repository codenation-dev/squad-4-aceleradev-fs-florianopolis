package models

import (
	"errors"
	"fmt"
)

var (
	// ErrCustNotFound - customer not found
	ErrCustNotFound = errors.New("Cliente não encontrado")
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
// TODO: será que posso usar esta função para receber os dados do arquivo.csv

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
	if len(customers) == 0 {
		return nil, ErrCustNotFound
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
	if len(customers) == 0 {
		return nil, ErrCustNotFound
	}
	return customers, nil
}

// GetCustomersByName returns a bank customer
func GetCustomersByName(name string) ([]Customer, error) {
	db := Connect()
	defer db.Close()
	sql := "SELECT * FROM customers WHERE name = $1"
	rows, err := db.Query(sql, name)
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
	if len(customers) == 0 {
		return nil, ErrCustNotFound
	}
	return customers, nil
}

// UpdateCustomer updates a customer
func UpdateCustomer(c Customer) (int64, error) {
	db := Connect()
	defer db.Close()
	sql := `UPDATE customers 
			SET name = $1, wage = $2, is_public = $3, sent_warning = $4
			WHERE cid = $5`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	rows, err := stmt.Exec(c.Name, c.Wage, c.IsPublic, c.SentWarning, c.CID)
	if err != nil {
		return 0, err
	}
	return rows.RowsAffected()
}
