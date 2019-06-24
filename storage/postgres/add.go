package postgres

import "codenation/squad-4-aceleradev-fs-florianopolis/entity"

// AddCustomer inserts a new customer on the DB
func (s *Storage) AddCustomer(c entity.Customer) error {
	_, err := s.db.Exec(`INSERT INTO customers (name, wage, is_public, sent_warning)
						VALUES ($1, $2, $3, $4)`,
		&c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
	return err
}

// AddUser inserts a new customer on the DB
func (s *Storage) AddUser(u entity.User) error {
	_, err := s.db.Exec(`INSERT INTO users (login, email, pass)
						VALUES ($1, $2, $3)`,
		&u.Login, &u.Email, &u.Pass)
	return err
}
