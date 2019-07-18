package postgres

import (
	"fmt"
	"log"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

// CreateUser inserts a new user on the DB
func (s *Storage) CreateUser(u entity.User) error {

	_, err := s.db.Exec(`INSERT INTO users (email, password)
						VALUES ($1, $2)`,
		&u.Email, &u.Password)
	return err
}

// CreatePublicFunc inserts a new public agent on the DB
func (s *Storage) CreatePublicFunc(tableName string, pp ...entity.PublicFunc) error {

	var query = fmt.Sprintf(`INSERT INTO %s (complete_name, short_name, wage, departament, function) VALUES `, tableName)

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
			query = fmt.Sprintf(`INSERT INTO %s (complete_name, short_name, wage, departament, function) VALUES `, tableName)
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
func (s *Storage) CreateCustomer(tableName string, cc ...entity.Customer) error {

	var query = fmt.Sprintf(`INSERT INTO %s (name) VALUES `, tableName)

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
			query = fmt.Sprintf(`INSERT INTO %s (name) VALUES `, tableName)
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
