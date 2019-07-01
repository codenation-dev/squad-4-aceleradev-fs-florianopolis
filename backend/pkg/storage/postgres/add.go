package postgres

import (
	"fmt"
	"log"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/utils"
)

// AddCustomer inserts a new customer on the DB
func (s *Storage) AddCustomer(c entity.Customer) error {
	_, err := s.db.Exec(`INSERT INTO customers (name, wage, is_public, sent_warning)
						VALUES ($1, $2, $3, $4)`,
		&c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
	return err
}

// AddUser inserts a new user on the DB
func (s *Storage) AddUser(u entity.User) error {
	bPass, err := utils.Bcrypt(u.Pass)
	if err != nil {
		return err
	}
	u.Pass = string(bPass)
	_, err = s.db.Exec(`INSERT INTO users (login, email, pass)
						VALUES ($1, $2, $3)`,
		&u.Login, &u.Email, &u.Pass)
	return err
}

// AddWarning inserts a new warning on the DB
func (s *Storage) AddWarning(w entity.Warning) error {
	_, err := s.db.Exec(`INSERT INTO warnings (dt, msg, sent_to, from_customer)
						VALUES ($1, $2, $3, $4)`,
		&w.Dt, &w.Message, &w.SentTo, &w.FromCustomer)
	return err
}

// AddPublicFunc inserts a new public agent on the DB
func (s *Storage) AddPublicFunc(pp ...entity.PublicFunc) error {
	var query = `INSERT INTO public_funcs (name, wage, place) VALUES `
	var vals = []interface{}{}
	batch := 20000
	i := 0
	for _, p := range pp {
		inc := fmt.Sprintf("($%v, $%v, $%v), ", (1 + (i * 3)), (2 + (i * 3)), (3 + (i * 3)))
		i++
		query += inc
		vals = append(vals, p.Name, p.Wage, p.Place)
		if i%batch == 0 && i != 0 {
			q := query[0 : len(query)-2]
			_, err := s.db.Exec(q, vals...)
			if err != nil {
				log.Fatalf("Erro na função AddPublicFunc 1", err)
			}
			query = `INSERT INTO public_funcs (name, wage, place) VALUES `
			vals = []interface{}{}
			i = 0
		}
	}
	q := query[0 : len(query)-2]

	_, err := s.db.Exec(q, vals...)
	if err != nil {
		log.Fatalf("Erro na função AddPublicFunc 2", err)
	}

	return err
}
