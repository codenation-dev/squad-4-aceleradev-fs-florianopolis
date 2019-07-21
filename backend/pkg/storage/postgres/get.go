package postgres

import (
	"fmt"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
)

// ByID

// GetCustomerByID read a customer from the DB, given the id
func (s *Storage) GetCustomerByID(id int) (model.Customer, error) {
	c := model.Customer{}
	query := "SELECT * FROM customers WHERE id=$1"
	err := s.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
	if err != nil {
		return model.Customer{}, err
	}
	return c, err
}

// GetUserByID read a customer from the DB, given the id
func (s *Storage) GetUserByID(id int) (model.User, error) {
	u := model.User{}
	query := "SELECT * FROM users WHERE id=$1"
	err := s.db.QueryRow(query, id).Scan(&u.ID, &u.Email, &u.Pass)
	if err != nil {
		return model.User{}, err
	}
	return u, err
}

// GetWarningByID read a warning from the DB, given the id
func (s *Storage) GetWarningByID(id int) (model.Warning, error) {
	w := model.Warning{}
	query := "SELECT * FROM warnings WHERE id=$1"
	err := s.db.QueryRow(query, id).Scan(&w.ID, &w.Dt, &w.Message, &w.FromCustomer, &w.SentTo)
	if err != nil {
		return model.Warning{}, err
	}
	return w, err
}

// GetPublicByID read a public_func from the DB, given the id
func (s *Storage) GetPublicByID(id int) (model.PublicFunc, error) {
	p := model.PublicFunc{}
	query := "SELECT * FROM public_funcs WHERE id=$1"
	err := s.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Wage, &p.Place)
	if err != nil {
		return model.PublicFunc{}, err
	}
	return p, err
}

//All

// GetAllCustomers return all customers from the DB
func (s *Storage) GetAllCustomers() ([]model.Customer, error) {
	customers := []model.Customer{}
	rows, err := s.db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := model.Customer{}
		err := rows.Scan(&c.ID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// GetAllUsers return all customers from the DB
func (s *Storage) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.ID, &u.Email, &u.Pass)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// GetAllWarnings return all customers from the DB
func (s *Storage) GetAllWarnings() ([]model.Warning, error) {
	warnings := []model.Warning{}
	rows, err := s.db.Query("SELECT * FROM warnings")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		w := model.Warning{}
		err := rows.Scan(&w.ID, &w.Dt, &w.Message, &w.FromCustomer, &w.SentTo)
		if err != nil {
			return nil, err
		}
		warnings = append(warnings, w)
	}
	return warnings, nil
}

// ByName

// GetCustomerByName returns a slice of customers with the given pattern in the name column
func (s *Storage) GetCustomerByName(pattern string) ([]model.Customer, error) {
	customers := []model.Customer{}
	// query := fmt.Sprintf("SELECT * FROM customers WHERE name LIKE '%%%s%%'", pattern)
	query := `SELECT * FROM customers WHERE name LIKE $1`
	fmt.Println(pattern)
	// pat := fmt.Sprintf()
	// fmt.Println(pat)
	rows, err := s.db.Query(query, "%"+pattern+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := model.Customer{}
		err := rows.Scan(&c.ID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// GetUserByEmail returns a slice of customers with the given pattern in the name column
func (s *Storage) GetUserByEmail(pattern string) (model.User, error) {
	u := model.User{}
	query := "SELECT * FROM users WHERE email = $1"
	err := s.db.QueryRow(query, pattern).Scan(&u.ID, &u.Email, &u.Pass)
	// if err != nil {
	// 	return nil, err
	// }
	// for rows.Next() {
	// 	u := model.User{}
	// 	err := rows.Scan(&u.ID, &u.Login, &u.Email, &u.Pass)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, u)
	// }
	return u, err
}

// GetWarningByCustomer returns a slice of warnings with the given pattern in the sent_to column
func (s *Storage) GetWarningByCustomer(pattern string) ([]model.Warning, error) {
	warnings := []model.Warning{}
	query := fmt.Sprintf("SELECT * FROM warnings WHERE from_customer LIKE '%%%s%%'", pattern)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		w := model.Warning{}
		err := rows.Scan(&w.ID, &w.Dt, &w.Message, &w.SentTo, &w.FromCustomer)
		if err != nil {
			return nil, err
		}
		warnings = append(warnings, w)
	}
	return warnings, nil
}

// GetWarningByUser returns a slice of warnings with the given pattern in the sent_to column
func (s *Storage) GetWarningByUser(pattern string) ([]model.Warning, error) {
	warnings := []model.Warning{}
	query := fmt.Sprintf("SELECT * FROM warnings WHERE sent_to LIKE '%%%s%%'", pattern)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		w := model.Warning{}
		err := rows.Scan(&w.ID, &w.Dt, &w.Message, &w.SentTo, &w.FromCustomer)
		if err != nil {
			return nil, err
		}
		warnings = append(warnings, w)
	}
	return warnings, nil
}

// GetPublicByWage returns a slice of public agents that earns more than the given pattern
func (s *Storage) GetPublicByWage(pattern float32) ([]model.PublicFunc, error) {
	publicFuncs := []model.PublicFunc{}
	query := `SELECT * FROM public_funcs WHERE wage > $1`
	rows, err := s.db.Query(query, pattern)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pf := model.PublicFunc{}
		err := rows.Scan(&pf.ID, &pf.Name, &pf.Wage, &pf.Place)
		if err != nil {
			return nil, err
		}
		publicFuncs = append(publicFuncs, pf)
	}
	return publicFuncs, nil
}
