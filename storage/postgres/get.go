package postgres

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
	"fmt"
)

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

// GetUserByID read a customer from the DB, given the id
func (s *Storage) GetUserByID(id int) (entity.User, error) {
	u := entity.User{}
	query := "SELECT * FROM users WHERE id=$1"
	err := s.db.QueryRow(query, id).Scan(&u.ID, &u.Login, &u.Email, &u.Pass)
	if err != nil {
		return entity.User{}, err
	}
	return u, err
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

// GetAllUsers return all customers from the DB
func (s *Storage) GetAllUsers() ([]entity.User, error) {
	users := []entity.User{}
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := entity.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Email, &u.Pass)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
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

// GetUserByEmail returns a slice of customers with the given pattern in the name column
func (s *Storage) GetUserByEmail(pattern string) ([]entity.User, error) {
	users := []entity.User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE email LIKE '%%%s%%'", pattern)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := entity.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Email, &u.Pass)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
