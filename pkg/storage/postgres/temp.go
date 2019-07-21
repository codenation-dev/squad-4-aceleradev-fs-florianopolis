package publicFunc

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

const (
	// DB data
	DRIVER_NAME = "postgres"
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	// DB_NAME     = "uati"
	SSLMODE = "disable"
	HOST    = "172.17.0.2"
	PORT    = "5432"
)

// Storage stores data ia a postgresql db
type Storage struct {
	db *sql.DB
}

// Connect implements the connection to the db
func Connect(dbName string) (*sql.DB, error) {
	connString := fmt.Sprintf(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, HOST, PORT, dbName, SSLMODE,
	))
	return sql.Open("postgres", connString)
}

// NewStorage creates a new instance of Storage
func NewStorage(dbName string) *Storage {
	db, err := Connect(dbName)
	if err != nil {
		log.Fatal(err)
	}
	s := new(Storage)
	s.db = db
	return s
}

// CreatePublicFunc inserts a new public agent on the DB
func (s *Storage) CreatePublicFunc(pp ...PublicFunc) error {

	var query = `INSERT INTO public_func (complete_name, short_name, wage, departament, function) VALUES `

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

func (filter Filter) makeFilter(paginated bool) string {
	where := "where 1=1"

	if filter.ID != 0 {
		where += fmt.Sprintf(" AND id_funcionario = %d", filter.ID)
	}
	if filter.Nome != "" {
		where += fmt.Sprintf(" AND nome ILIKE '%%%s%%'", filter.Nome)
	}
	if filter.Cargo != "" {
		where += fmt.Sprintf(" AND cargo ILIKE '%%%s%%'", filter.Cargo)
	}
	if filter.Orgao != "" {
		where += fmt.Sprintf(" AND orgao ILIKE '%%%s%%'", filter.Orgao)
	}

	if paginated {
		where += " ORDER BY " + filter.SortBy
		fmt.Println(filter)
		if filter.Desc {
			where += " desc"
		} else {
			where += " asc"
		}
		where += ` limit ` + strconv.FormatInt(filter.Offset, 10)
		where += ` offset ` + strconv.FormatInt(filter.Page*filter.Offset, 10)
	}

	return where
}

// ReadPublicFunc returns a slice with all public agents
func (s *Storage) ReadPublicFunc(filter Filter) ([]PublicFunc, error) {
	query := `SELECT  complete_name, short_name, wage, departament, function FROM public_func`
	query += filter.makeFilter(true)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	return scanRowsPublicFunc(rows)

	// if err != nil {
	// 	if err.Error() == `pq: relation "public_func_sp_2019_abril" does not exist` {
	// 		err = s.fetchPublicFuncData(uf, year, month)
	// 		fmt.Println("g***************")
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		return s.ReadAllPublicFunc(uf, year, month)
	// 	} else {

	// 		return nil, err
	// 	}
	// } else {
	// 	var count int
	// 	_ = s.db.QueryRow(fmt.Sprintf("select count (*) from %s", tableName)).Scan(&count)
	// 	if count == 0 {
	// 		err = s.fetchPublicFuncData(uf, year, month)
	// 		fmt.Println("g***************")
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		return s.ReadAllPublicFunc(uf, year, month)
	// 	}
	// }

}

func scanRowsPublicFunc(rows *sql.Rows) ([]PublicFunc, error) {
	publicFuncs := []PublicFunc{}

	for rows.Next() {
		pf := PublicFunc{}
		err := rows.Scan(&pf.CompleteName, &pf.ShortName, &pf.Wage, &pf.Departament, &pf.Function)
		if err != nil {
			return nil, err
		}
		publicFuncs = append(publicFuncs, pf)
	}
	return publicFuncs, nil
}

func (s *Storage) ImportPublicFunc(month, year string) error {
	_, err := s.db.Exec("DROP TABLE public_func")
	if err != nil {
		log.Fatal("drop table", err)
	}
	_, err = s.db.Exec("")
	return err
}
