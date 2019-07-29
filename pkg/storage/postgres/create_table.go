package postgres

import (
	"fmt"
)

// CreatePublicFuncTable inserts a new table if it do not exists and import the file
func (s *Storage) createPublicFuncTable(tableName string) error {
	// tableName := fmt.Sprintf("public_func_%s_%s_%s", uf, year, month)
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL primary key,
		complete_name VARCHAR(255),
		short_name VARCHAR(30),
		wage NUMERIC(15,2),
		departament VARCHAR(100),
		function VARCHAR(100),
		relevancia smallint
		)`, tableName)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf(`CREATE INDEX idx_short_name ON %s (short_name)`, tableName)
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateCustomerTable inserts a new table if it do not exists and import the file
func (s *Storage) createCustomerTable(company string) error {
	company = "customer"
	query := fmt.Sprintf(`CREATE TABLE %s (id SERIAL primary key, name VARCHAR(30))`, company)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`CREATE INDEX idx_nome_customer ON %s (name)`, company)
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) createUsersTable(name string) error {
	name = "users"
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL,
		email VARCHAR(50) UNIQUE,
		password VARCHAR(100))`, name)
	fmt.Println(query)
	_, err := s.db.Exec(query)
	return err
}

func (s *Storage) dropTable(name string) error {
	_, err := s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", name))
	return err
}
