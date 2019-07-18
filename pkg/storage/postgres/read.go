package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/importing"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

func (s *Storage) fetchPublicFuncData(uf, year, month string) error {
	tableName := fmt.Sprintf("public_func_%s_%s_%s", uf, year, month)
	err := s.createPublicFuncTable(tableName)
	if err != nil {
		return err
	}
	publicFuncs, err := importing.FetchPublicAgentsFile(uf, month, year)
	if err != nil {
		return err
	}
	err = s.CreatePublicFunc(tableName, publicFuncs...)
	if err != nil {
		return err
	}
	return nil
}

// ReadAllPublicFunc returns a slice with all public agents
func (s *Storage) ReadAllPublicFunc(uf, year, month string) ([]entity.PublicFunc, error) {
	tableName := fmt.Sprintf("public_func_%s_%s_%s", uf, year, month)

	query := fmt.Sprintf(`SELECT  complete_name, short_name, wage, departament, function FROM %s`, tableName)
	rows, err := s.db.Query(query)
	if err != nil {
		if err.Error() == `pq: relation "public_func_sp_2019_abril" does not exist` {
			err = s.fetchPublicFuncData(uf, year, month)
			if err != nil {
				panic(err)
			}
			return s.ReadAllPublicFunc(uf, year, month)
		} else {
			return nil, err
		}
	}
	return scanRowsPublicFunc(rows)
}

func (s *Storage) fetchCustomerData(company string) error {
	var err error
	err = s.createCustomerTable(company)
	if err != nil {
		return err
	}
	customers, err := importing.FetchCustomerData(company)
	if err != nil {
		return err
	}
	return s.CreateCustomer(company, customers...)
}

// ReadAllCustomers return all customers from the DB
func (s *Storage) ReadAllCustomers(company string) ([]entity.Customer, error) {
	customers := []entity.Customer{}
	query := fmt.Sprintf("SELECT name FROM %s", company)
	rows, err := s.db.Query(query)
	if err != nil {
		if err.Error() == fmt.Sprintf(`pq: relation "%s" does not exist`, company) {
			err := s.fetchCustomerData(company)
			if err != nil {
				panic(err)
			}
			return s.ReadAllCustomers(company)
		}
		return nil, err
	}
	for rows.Next() {
		c := entity.Customer{}
		err := rows.Scan(&c.Name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func scanRowsPublicFunc(rows *sql.Rows) ([]entity.PublicFunc, error) {
	publicFuncs := []entity.PublicFunc{}

	for rows.Next() {
		pf := entity.PublicFunc{}
		err := rows.Scan(&pf.CompleteName, &pf.ShortName, &pf.Wage, &pf.Departament, &pf.Function)
		if err != nil {
			return nil, err
		}
		publicFuncs = append(publicFuncs, pf)
	}
	return publicFuncs, nil
}

func (s *Storage) readPublicFuncByList(funcTableName, field string, list []interface{}) ([]entity.PublicFunc, error) {
	query := fmt.Sprintf(`SELECT complete_name, short_name, wage, departament, function 
						FROM %s WHERE %s IN (`, funcTableName, field)
	for n := 1; n < len(list)+1; n++ {
		query += fmt.Sprintf("$%s, ", strconv.Itoa(n))
	}
	query = query[:len(query)-2] + ")"

	rows, err := s.db.Query(query, list...)
	if err != nil {
		return nil, err
	}
	return scanRowsPublicFunc(rows)
}

// CompareCustomerPublicFunc returns a slice with all public agents that already are bank's customers
func (s *Storage) CompareCustomerPublicFunc(funcTableName, customerTableName string) ([]entity.PublicFunc, error) {
	var err error
	customers := []entity.Customer{}

	customers, err = s.ReadAllCustomers(customerTableName)
	if err != nil {
		return nil, err
	}
	names := []interface{}{}
	for _, c := range customers {
		names = append(names, c.Name)
	}
	return s.readPublicFuncByList(funcTableName, "short_name", names)
}
