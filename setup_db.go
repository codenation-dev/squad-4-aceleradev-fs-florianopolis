package main

import (
	"fmt"
)

const publicTableCreationQuery = `CREATE TABLE IF NOT EXISTS public_employees
	(
		id SERIAL,
		name TEXT NOT NULL,
		wage NUMERIC(10,2) DEFAULT 0.00,
		local TEXT NOT NULL
	)`

const customersTableCreationQuery = `CREATE TABLE IF NOT EXISTS customers
	(
		id SERIAL,
		name TEXT NOT NULL,
		wage NUMERIC(10,2) DEFAULT 0.00,
		isPublic bit DEFAULT NULL,
		sentWarning TEXT DEFAULT NULL
	)`

const usersTableCreationQuery = `CREATE TABLE IF NOT EXISTS users
	(
		id SERIAL,
		email TEXT NOT NULL,
		password TEXT
	)`

const warningsTableCreationQuery = `CREATE TABLE IF NOT EXISTS warnings
	(
		id SERIAL,
		dt TIMESTAMP,
		message TEXT,
		sentTo text,
		fromCustomer TEXT
	)`

func dropTable(table string) error {
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("could not openDB: %v", err)
	}
	defer db.Close()

	if _, err := db.Exec(fmt.Sprintf("DROP TABLE %s", table)); err != nil {
		return fmt.Errorf("Could not drop %s table: %v", table, err)
	}
	return nil
}

func createAllTables() error {
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("could not openDB: %v", err)
	}
	defer db.Close()

	if _, err := db.Exec(customersTableCreationQuery); err != nil {
		return fmt.Errorf("Could not create customer table: %v", err)
	}
	if _, err := db.Exec(usersTableCreationQuery); err != nil {
		return fmt.Errorf("Could not create users table: %v", err)
	}
	if _, err := db.Exec(warningsTableCreationQuery); err != nil {
		return fmt.Errorf("Could not create warnings table: %v", err)
	}
	if _, err := db.Exec(publicTableCreationQuery); err != nil {
		return fmt.Errorf("Could not create public agents table: %v", err)
	}
	return nil
}

const listTablesQuery = `SELECT table_schema || '.' || table_name
						FROM information_schema.tables
						WHERE table_type = 'BASE TABLE'
						AND table_schema NOT IN ('pg_catalog', 'information_schema')`

func listTables() ([]string, error) {
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("could not opneDB(): %v", err)
	}
	defer db.Close()

	rows, err := db.Query(listTablesQuery)
	if err != nil {
		return nil, fmt.Errorf("Could not listTables: %v", err)
	}
	var tables []string
	var table string
	for rows.Next() {
		err = rows.Scan(&table)
		if err != nil {
			return nil, fmt.Errorf("could not rows.Scan the listTables: %v", err)
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func deleteAllRows(table string) (int64, error) {
	db, err := openDB()
	if err != nil {
		return 0, fmt.Errorf("could not openDB(): %v", err)
	}
	defer db.Close()

	result, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
	if err != nil {
		return 0, fmt.Errorf("could not deleteAllRows(): %v", err)
	}
	return result.RowsAffected()
}
