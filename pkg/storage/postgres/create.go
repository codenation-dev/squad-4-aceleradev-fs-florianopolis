package postgres

import (
	"fmt"
	"log"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/importing"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

// CreateUser inserts a new user on the DB
func (s *Storage) CreateUser(u entity.User) error {

	_, err := s.db.Exec(`INSERT INTO users (email, password)
						VALUES ($1, $2)`,
		&u.Email, &u.Password)
	return err
}

// ImportPublicFunc implement the import routine
// it drops the old table, and populate it again with new values
func (s *Storage) ImportPublicFunc(month, year string) error {
	_, err := s.db.Exec("DROP TABLE IF EXISTS  public_func")
	if err != nil {
		log.Fatal("drop table", err)
	}
	_, err = s.db.Exec(`CREATE TABLE public_func (
		id SERIAL,
		name VARCHAR(100),
		wage NUMERIC(10,2),
		departament VARCHAR(50),
		function VARCHAR(50)
		)`)
	return err

	publicFuncs, err := importing.ImportPublicFuncFile(month, year)
	if err != nil {
		return err
	}

	return s.CreatePublicFunc(publicFuncs...)
}

// CreatePublicFunc inserts a new public agent on the DB
func (s *Storage) CreatePublicFunc(pp ...entity.PublicFunc) error {

	var query = `INSERT INTO %s (complete_name, short_name, wage, departament, function) VALUES `

	var vals = []interface{}{}
	i := 0
	numberOfFields := 5
	batch := 13000 // (65000 / numberOfFields )
	for _, p := range pp {
		// we cannot use the '?' in the query because limitations of the driver
		// so we used the '$1, $2, $3...' notation
		multiply := i * numberOfFields
		inc := fmt.Sprintf("($%v, $%v, $%v, $%v, $%v), ", (1 + multiply), (2 + multiply), (3 + multiply), (4 + multiply), (5 + multiply))
		i++ // can't use the built in index because we restart it to make the batches
		query += inc
		vals = append(vals, p.CompleteName, p.ShortName, p.Wage, p.Departament, p.Function)
		if i%batch == 0 && i != 0 {
			q := query[0 : len(query)-2]

			_, err := s.db.Exec(q, vals...)
			if err != nil {
				log.Fatalf("error executing batch (%v)", err)
			}
			// restart the vars to a new batch
			query = `INSERT INTO public_func (complete_name, short_name, wage, departament, function) VALUES `
			vals = []interface{}{}
			i = 0

		}
	}
	q := query[0 : len(query)-2]

	_, err := s.db.Exec(q, vals...)
	if err != nil {
		log.Fatalf("error executing the remaining batch (%v)", err)
	}

	return err
}

// CreateCustomer inserts a new customer on the DB
func (s *Storage) CreateCustomer(cc ...entity.Customer) error {

	var query = `INSERT INTO customer (name) VALUES `

	var vals = []interface{}{}
	batch := 60000 // if more fields in data, we have to change the batch
	i := 0
	for _, c := range cc {
		// we cannot use the '?' in the query because limitations of the driver
		// so we used the '$1, $2, $3...' notation
		inc := fmt.Sprintf("($%v), ", (1 + i))
		i++ // can't use the built in index because we restart it to make the batches
		query += inc
		vals = append(vals, c.Name)
		if i%batch == 0 && i != 0 {
			q := query[0 : len(query)-2]

			_, err := s.db.Exec(q, vals...)
			if err != nil {
				return fmt.Errorf("error executing batch (%v)", err)
			}
			// restart the vars to a new batch
			query = `INSERT INTO customer (name) VALUES `
			vals = []interface{}{}
			i = 0

		}
	}
	q := query[0 : len(query)-2]

	_, err := s.db.Exec(q, vals...)
	if err != nil {
		return fmt.Errorf("error executing the remaining batch (%v)", err)
	}

	return err
}

// // AddCustomer inserts a new customer on the DB
// func (s *Storage) AddCustomer(c model.Customer) error {
// 	_, err := s.db.Exec(`INSERT INTO customers (name, wage, is_public, sent_warning)
// 						VALUES ($1, $2, $3, $4)`,
// 		&c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
// 	fmt.Println(c, err)

// 	return err
// }
