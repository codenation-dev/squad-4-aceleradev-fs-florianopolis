package postgres

// DeleteCustomerByID delets a customer from the db
func (s *Storage) DeleteCustomerByID(id int) error {
	_, err := s.db.Exec(`DELETE FROM customers WHERE id=$1`, id)
	return err
}

// DeleteUserByID delets a customer from the db
func (s *Storage) DeleteUserByID(id int) error {
	_, err := s.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
