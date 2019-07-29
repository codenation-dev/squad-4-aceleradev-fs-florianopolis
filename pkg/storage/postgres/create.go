package postgres

import (
	"fmt"
	"log"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/importing"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

// CreateUser inserts a new user on the DB
func (s *Storage) CreateUser(u entity.User) error {
	err := s.createUsersTable("users")
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`INSERT INTO users (email, password)
						VALUES ($1, $2)`,
		&u.Email, &u.Password)
	return err
}

func (s *Storage) ImportCustomer() error {
	_, err := s.db.Exec("DROP TABLE IF EXISTS customer")

	if err != nil {
		return err
	}
	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS customer (
		id SERIAL,
		name VARCHAR(30))`)
	if err != nil {
		return err
	}

	customers, err := importing.ImportCustomer()
	if err != nil {
		return err
	}

	return s.CreateCustomer(customers...)
}

// ImportPublicFunc implement the import routine
// it drops the old table, and populate it again with new values
func (s *Storage) ImportPublicFunc(month, year string) error {
	err := s.dropTable("public_func")
	if err != nil {
		return err
	}

	err = s.createPublicFuncTable("public_func")
	if err != nil {
		return err
	}

	publicFuncs, err := importing.ImportPublicFuncFile(month, year)
	if err != nil {
		return err
	}

	return s.CreatePublicFunc(publicFuncs...)
}

// CreatePublicFunc inserts a new public agent on the DB
func (s *Storage) CreatePublicFunc(pp ...entity.PublicFunc) error {

	var query = `INSERT INTO public_func (complete_name, short_name, wage, departament, function, relevancia) VALUES `

	var vals = []interface{}{}
	i := 0
	numberOfFields := 6
	batch := 10800 // (65000 / numberOfFields )
	for _, p := range pp {
		// we cannot use the '?' in the query because limitations of the driver
		// so we used the '$1, $2, $3...' notation
		multiply := i * numberOfFields
		inc := fmt.Sprintf("($%v, $%v, $%v, $%v, $%v, $%v), ", (1 + multiply), (2 + multiply), (3 + multiply),
			(4 + multiply), (5 + multiply), (6 + multiply))
		i++ // can't use the built in index because we restart it to make the batches
		query += inc
		vals = append(vals, p.CompleteName, p.ShortName, p.Wage, p.Departament, p.Function, p.Relevancia)
		if i%batch == 0 && i != 0 {
			q := query[0 : len(query)-2]

			_, err := s.db.Exec(q, vals...)
			if err != nil {
				log.Println(vals[32390:32411])
				return fmt.Errorf("error executing batch (%v)", err)
			}
			// restart the vars to a new batch
			query = `INSERT INTO public_func (complete_name, short_name, wage, departament, function, relevancia) VALUES `
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
