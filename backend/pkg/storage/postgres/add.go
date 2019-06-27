package postgres

import "codenation/squad-4-aceleradev-fs-florianopolis/pkg/entity"

// AddCustomer inserts a new customer on the DB
func (s *Storage) AddCustomer(c entity.Customer) error {
	_, err := s.db.Exec(`INSERT INTO customers (name, wage, is_public, sent_warning)
						VALUES ($1, $2, $3, $4)`,
		&c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
	return err
}

// AddUser inserts a new user on the DB
func (s *Storage) AddUser(u entity.User) error {
	_, err := s.db.Exec(`INSERT INTO users (login, email, pass)
						VALUES ($1, $2, $3)`,
		&u.Login, &u.Email, &u.Pass)
	return err
}

// AddWarning inserts a new warning on the DB
func (s *Storage) AddWarning(w entity.Warning) error {
	_, err := s.db.Exec(`INSERT INTO warnings (dt, message, sent_to, from_customer)
						VALUES ($1, $2, $3, $4)`,
		&w.Dt, &w.Message, &w.SentTo, &w.FromCustomer)
	return err
}

// AddPublicFunc inserts a new public agent on the DB
func (s *Storage) AddPublicFunc(p entity.PublicFunc) error {
	_, err := s.db.Exec(`INSERT INTO public_funcs (name, wage, place)
						VALUES ($1, $2, $3)`, &p.Name, &p.Wage, &p.Place)
	return err
}

// func (s *Storage) LoadPublicFuncFile() error {
// 	_, err := s.db.Exec(`LOAD DATA INFILE 'Remuneracao.txt' INTO TABLE public_funcs
// 						FIELDS TERMINATED BY ';'`)
// 	return err
// }
