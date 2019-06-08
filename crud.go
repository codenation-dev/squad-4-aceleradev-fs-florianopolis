package main

import (
	"database/sql"
	"fmt"
	"log"
)

func openDB() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, fmt.Errorf("Could not sql.Open: %v", err)
	}
	return db, nil
}

func read(query string) (*sql.Rows, error) {
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("could not createCustomer(): %v", err)
	}
	defer db.Close()
	return db.Query(query)
}

// Return just one row
// "by" is the rest of the query Ex: " WHERE name = 'Fulano'"
func (pe *publicEmployee) read(by string) {
	// "by" is the rest of the query Ex: " WHERE name = 'Fulano'"

	// TODO: Fazer tipo este método para todas as tabelas e slices e usar os
	// mesmos nomes de função para fazer uma interface e reunir tudo facinho

	db, _ := openDB()
	q := fmt.Sprintf("SELECT * FROM public_employees %s", by)
	rows := db.QueryRow(q)
	err := rows.Scan(&pe.ID, &pe.Name, &pe.Wage, &pe.Local)
	if err != nil {
		log.Fatalf("could not use method publicEmployee.read(): %v", err)
	}
}

// func create() {
// 	db, err := openDB()
// 	if err != nil {
// 		return "", fmt.Errorf("could not createCustomer(): %v", err)
// 	}
// 	defer db.Close()
// }

func createCustomer(c customer) (string, error) {
	db, err := openDB()
	if err != nil {
		return "", fmt.Errorf("could not createCustomer(): %v", err)
	}
	defer db.Close()

	var lastInsertedName string
	err = db.QueryRow(
		`INSERT INTO customers (name, wage, isPublic, sentWarning)
		VALUES($1, $2, $3, $4) returning name;`, c.Name, c.Wage, c.IsPublic, c.SentWarning).
		Scan(&lastInsertedName)
	if err != nil {
		return "", fmt.Errorf("could not addCustomer(): %v", err)
	}
	return lastInsertedName, nil
}

func createEmployee(e publicEmployee) (string, error) {
	db, err := openDB()
	if err != nil {
		return "", fmt.Errorf("could not createEmployee(): %v", err)
	}
	defer db.Close()

	var lastInsertedName string
	err = db.QueryRow(
		`INSERT INTO public_employees (name, wage, local)
		VALUES($1, $2, $3) returning name;`, e.Name, e.Wage, e.Local).
		Scan(&lastInsertedName)
	if err != nil {
		return "", fmt.Errorf("could not addCustomer(): %v", err)
	}
	return lastInsertedName, nil
}

// func readDB(query string) *sql.Rows, err {
// 	db, err := openDB()
// 	if err != nil {
// 		return nil, fmt.Errorf("could openDB(): %v", err)
// 	}
// 	defer db.Close()

// 	var e publicEmployee
// 	var ee []publicEmployee
// 	// q := fmt.Sprintf("SELECT * FROM public_employees %s", by)
// 	rows, err := db.Query(q)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not read public_employee table: %v", err)
// 	}
// }

// "by" is the rest of the query Ex: " WHERE name = 'Fulano'"
func readEmployee(by string) ([]publicEmployee, error) {
	// db, err := openDB()
	// if err != nil {
	// 	return nil, fmt.Errorf("could readAllCustomers(): %v", err)
	// }
	// defer db.Close()

	var e publicEmployee
	var ee []publicEmployee
	q := fmt.Sprintf("SELECT * FROM public_employees %s", by)
	rows, err := read(q)
	if err != nil {
		return nil, fmt.Errorf("could not read public_employee table: %v", err)
	}
	for rows.Next() {
		err := rows.Scan(&e.ID, &e.Name, &e.Wage, &e.Local)
		if err != nil {
			return nil, fmt.Errorf("could not row.Scan(): %v", err)
		}
		ee = append(ee, e)
	}
	return ee, nil
}

// "by" is the rest of the query Ex: " WHERE name = 'Fulano'"
func readCustomers(by string) ([]customer, error) {
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("could not readAllCustomers(): %v", err)
	}
	defer db.Close()

	q := fmt.Sprintf("SELECT * FROM customers &s", by)
	rows, err := db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("could not db.Query(): %v", err)
	}

	var c customer
	var customers []customer
	for rows.Next() {
		rows.Scan(&c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		customers = append(customers, c)
	}
	return customers, nil
}
