package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
)

type CountInfo struct {
	Total int     `json:"total"`
	Media float64 `json:"media"`
	Maior float64 `json:"maior"`
	Menor float64 `json:"menor"`
}

func makeFuncFilter(filter reading.FuncFilter, paginated bool) string {
	var where string

	switch filter.Customer {
	case "yes":
		where = ` INNER JOIN customer ON public_func.short_name = customer.name`
	case "no":
		where = ` LEFT JOIN customer ON public_func.short_name = customer.name`
	}

	where += " where 1 = 1"

	if filter.ID != 0 {
		where += fmt.Sprintf(" AND id_funcionario = %d", filter.ID)
	}
	if filter.Nome != "" {
		where += fmt.Sprintf(" AND complete_name ILIKE '%%%s%%'", filter.Nome)
	}
	if filter.Cargo != "" {
		where += fmt.Sprintf(" AND function ILIKE '%%%s%%'", filter.Cargo)
	}
	if filter.Orgao != "" {
		where += fmt.Sprintf(" AND departament ILIKE '%%%s%%'", filter.Orgao)
	}
	if filter.Salario > 0 {
		where += fmt.Sprintf(" AND wage > %d", filter.Salario)
	}
	if filter.Relevancia > 0 {
		where += fmt.Sprintf(" AND relevancia = %d", filter.Relevancia)
	}

	if paginated {
		where += " ORDER BY " + filter.SortBy
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

func getCountStats(s *Storage, filter reading.FuncFilter) interface{} {
	queryCount := `SELECT COUNT (1) total, AVG (wage) media, MAX (wage) maior, MIN (wage) menor FROM public_func `
	queryCount += makeFuncFilter(filter, false)

	count := CountInfo{}
	err := s.db.QueryRow(queryCount).Scan(&count.Total, &count.Media, &count.Maior, &count.Menor)
	if err != nil {
		log.Println(err)
		return count
	}

	return count
}

// ReadPublicFunc returns a slice with all public agents
func (s *Storage) ReadPublicFunc(filter reading.FuncFilter) (interface{}, error) {
	query := `SELECT  complete_name, short_name, wage, departament, function, relevancia FROM public_func `
	query += makeFuncFilter(filter, true)

	rows, err := s.db.Query(query)
	if err != nil {
		err := s.ImportPublicFunc("Maio", "2019")
		if err != nil {
			log.Fatal(err)
		}
		err = s.ImportCustomer()
		if err != nil {
			log.Fatal(err)
		}
	}

	publicfuncs, err := scanRowsPublicFunc(rows)

	resp := map[string]interface{}{}
	resp["stats"] = getCountStats(s, filter)
	resp["list"] = publicfuncs

	return resp, nil
}

// StatsPublicFunc returns a slice with some stats
func (s *Storage) StatsPublicFunc(filter reading.FuncFilter) (interface{}, error) {
	query := `select concat(floor(wage/10000) + 1, '0k'), avg(wage), count(*) as qtd from public_func `
	query += makeFuncFilter(filter, false)

	query += " group by floor(wage/10000) ORDER BY floor(wage/10000)"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	stats := []entity.PublicStats{}
	for rows.Next() {
		s := entity.PublicStats{}
		err := rows.Scan(&s.Floor, &s.Avg, &s.Qtd)
		if err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	resp := map[string]interface{}{}
	resp["stats"] = stats
	resp["info"] = getCountStats(s, filter)

	return resp, nil
}

// DistPublicFunc returns a slice with some stats
func (s *Storage) DistPublicFunc(filter reading.FuncFilter) (map[string][]entity.PublicStats, error) {
	query := `select function, min(wage), max(wage) as maximo, avg(wage) as media, count(1) as qtd from public_func`
	query += makeFuncFilter(filter, false)
	query += " group by function order by qtd desc limit 20"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	stats := []entity.PublicStats{}
	for rows.Next() {
		s := entity.PublicStats{}
		err := rows.Scan(&s.Cargo, &s.Min, &s.Max, &s.Avg, &s.Qtd)
		if err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	query2 := `select departament, min(wage), max(wage) as maximo, avg(wage) as media, count(1) as qtd from public_func`
	query2 += makeFuncFilter(filter, false)
	query2 += " group by departament order by qtd desc limit 20"
	rows, err = s.db.Query(query)
	if err != nil {
		return nil, err
	}
	stats2 := []entity.PublicStats{}
	rows, err = s.db.Query(query2)
	for rows.Next() {
		s := entity.PublicStats{}
		err := rows.Scan(&s.Departament, &s.Min, &s.Max, &s.Avg, &s.Qtd)
		if err != nil {
			return nil, err
		}
		stats2 = append(stats2, s)
	}

	result := map[string][]entity.PublicStats{}
	result["por_cargo"] = stats
	result["por_orgao"] = stats2
	return result, nil
}

func (s *Storage) Query(q, offset, page string) (interface{}, error) {

	switch q {
	case "count_by_departament":
		return s.countByDepartament(q, offset, page)
	case "min_max_avg_wage":
		return s.minMaxAvgWage()
		// case "best_wage"

	}

	return nil, errors.New("parametro 'q' ainda não implementado, contatar administrador do sistema")
}

func (s *Storage) minMaxAvgWage() (interface{}, error) {
	query := `SELECT MAX (wage), MIN (wage), AVG (wage) FROM public_func`
	type row struct {
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Mean float64 `json:"mean"`
	}
	r := row{}
	err := s.db.QueryRow(query).Scan(&r.Max, &r.Min, &r.Mean)
	return r, err
}

func (s *Storage) countByDepartament(q, offset, page string) ([]interface{}, error) {
	query := `SELECT COUNT (id) AS count, ROUND(avg(wage), 2) as media, min(wage) as minimo, max(wage) as maximo, departament 
					FROM public_func 
					GROUP BY departament 
					ORDER BY count DESC 
					LIMIT $1 OFFSET $2`
	type row struct {
		Count         int     `json:"count"`
		MediaSalarial float32 `json:"media"`
		MenorSalario  float32 `json:"minimo"`
		MaiorSalario  float32 `json:"maximo"`
		Departament   string  `json:"function"`
	}

	npage, _ := strconv.Atoi(page)
	noffset, _ := strconv.Atoi(offset)

	rows, err := s.db.Query(query, offset, (npage-1)*noffset)
	if err != nil {
		return nil, err
	}

	resp := []interface{}{}
	for rows.Next() {
		r := row{}
		err := rows.Scan(&r.Count, &r.MediaSalarial, &r.MenorSalario, &r.MaiorSalario, &r.Departament)
		if err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}
	return resp, nil
}

func makeCustomerFilter(filter reading.CustFilter, paginated bool) string {
	where := " where 1=1"

	if filter.ID != 0 {
		where += fmt.Sprintf(" AND id = %d", filter.ID)
	}
	if filter.Name != "" {
		where += fmt.Sprintf(" AND name ILIKE '%%%s%%'", filter.Name)
	}

	if paginated {
		where += " ORDER BY " + filter.SortBy + " "
		if filter.Desc {
			where += " desc "
		} else {
			where += " asc"
		}
		where += ` limit ` + strconv.FormatInt(filter.Offset, 10)
		where += ` offset ` + strconv.FormatInt(filter.Page*filter.Offset, 10)
	}
	return where
}

// ReadCustomer return customers from the DB
func (s *Storage) ReadCustomer(filter reading.CustFilter) ([]entity.Customer, error) {
	customers := []entity.Customer{}
	query := "SELECT name FROM customer" + makeCustomerFilter(filter, true)
	rows, err := s.db.Query(query)
	if err != nil {
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
		err := rows.Scan(&pf.CompleteName, &pf.ShortName, &pf.Wage, &pf.Departament, &pf.Function, &pf.Relevancia)
		if err != nil {
			return nil, err
		}
		publicFuncs = append(publicFuncs, pf)
	}
	return publicFuncs, nil
}

// // CompareCustomerPublicFunc returns a slice with all public agents that already are bank's customers
// func (s *Storage) CompareCustomerPublicFunc(funcTableName, customerTableName string) ([]entity.PublicFunc, error) {
// 	var err error
// 	customers := []entity.Customer{}

// 	customers, err = s.ReadAllCustomers(customerTableName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	names := []interface{}{}
// 	for _, c := range customers {
// 		names = append(names, c.Name)
// 	}
// 	return s.readPublicFuncByList(funcTableName, "short_name", names)
// }
