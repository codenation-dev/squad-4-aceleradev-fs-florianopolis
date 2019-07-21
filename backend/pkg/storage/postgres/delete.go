package postgres

// DeleteCustomerByID delets a customer from the db
func (s *Storage) DeleteCustomerByID(id int) (int64, error) {
	res, err := s.db.Exec(`DELETE FROM customers WHERE id=$1`, id)
	if err != nil {
		panic(err)
	}
	return res.RowsAffected()
}

// DeleteUserByID delets a customer from the db
func (s *Storage) DeleteUserByID(id int) (int64, error) {
	res, err := s.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		panic(err)
	}
	return res.RowsAffected()
}

// DeleteWarningByID delets a customer from the db
func (s *Storage) DeleteWarningByID(id int) (int64, error) {
	res, err := s.db.Exec(`DELETE FROM warnings WHERE id=$1`, id)
	if err != nil {
		panic(err)
	}
	return res.RowsAffected()
}

// DeletePublicByID delets a customer from the db
func (s *Storage) DeletePublicByID(id int) (int64, error) {
	res, err := s.db.Exec(`DELETE FROM public_funcs WHERE id=$1`, id)
	if err != nil {
		panic(err)
	}
	return res.RowsAffected()
}
