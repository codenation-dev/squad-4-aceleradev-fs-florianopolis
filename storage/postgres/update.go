package postgres

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
	"fmt"
)

// UpdateCustomer replace some data mantaining the same id
func (s *Storage) UpdateCustomer(c entity.Customer) error {
	_, err := s.db.Exec(`UPDATE customers
						SET name=$1, wage=$2, is_public=$3, sent_warning=$4
						WHERE id=$5`, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning, &c.ID)
	if err != nil {
		return fmt.Errorf("could not update the customer: %v", err)
	}
	return nil
}

// UpdateUser replace some data mantaining the same id
func (s *Storage) UpdateUser(u entity.User) error {
	_, err := s.db.Exec(`UPDATE users
						SET login=$1, email=$2, pass=$3
						WHERE id=$4`, &u.Login, &u.Email, &u.Pass, &u.ID)
	if err != nil {
		return fmt.Errorf("could not update the user: %v", err)
	}
	return nil
}
