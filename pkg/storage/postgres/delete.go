package postgres

import "github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"

// DeleteUser deletes a customer from the db
func (s *Storage) DeleteUser(email string) error {
	res, err := s.db.Exec(`DELETE FROM users WHERE email=$1`, email)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		err = entity.ErrUserNotFound
	}
	return err
}
