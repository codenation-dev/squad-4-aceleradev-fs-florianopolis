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

// DeleteWarningByID delets a customer from the db
func (s *Storage) DeleteWarningByID(id int) error {
	_, err := s.db.Exec(`DELETE FROM warnings WHERE id=$1`, id)
	return err
}

// DeletePublicByID delets a customer from the db
func (s *Storage) DeletePublicByID(id int) error {
	_, err := s.db.Exec(`DELETE FROM public_funcs WHERE id=$1`, id)
	return err
}