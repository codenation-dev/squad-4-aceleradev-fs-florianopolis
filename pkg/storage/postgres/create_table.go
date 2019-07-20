package postgres

import (
	"fmt"
	"log"
)

// CreatePublicFuncTable inserts a new table if it do not exists and import the file
func (s *Storage) createPublicFuncTable(tableName string) error {
	// tableName := fmt.Sprintf("public_func_%s_%s_%s", uf, year, month)
	query := fmt.Sprintf(`CREATE TABLE %s (
		id SERIAL,
		complete_name TEXT,
		short_name TEXT,
		wage NUMERIC(10,2),
		departament TEXT,
		function TEXT
		)`, tableName)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateCustomerTable inserts a new table if it do not exists and import the file
func (s *Storage) createCustomerTable(company string) error {
	query := fmt.Sprintf(`CREATE TABLE %s (id SERIAL, name TEXT)`, company)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) createUsersTable(name string) error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL,
		email TEXT UNIQUE)`, name)
	_, err := s.db.Exec(query)
	return err
}

func (s *Storage) dropTable(name string) {
	_, err := s.db.Exec(fmt.Sprintf("DROP TABLE %s", name))
	if err != nil {
		log.Fatal(err)
	}
}
