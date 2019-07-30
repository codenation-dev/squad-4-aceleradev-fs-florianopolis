package database

import (
	"fmt"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
)

type ClienteFilter struct {
	Offset int64  `schema:"offset"`
	Page   int64  `schema:"page"`
	SortBy string `schema:"sortby"`
	Desc   bool   `schema:"desc"`

	//User specific filters
	ID   int64  `schema:"id"`
	Nome string `schema:"nome"`
}

func (filter ClienteFilter) makeFilter(paginated bool) string {
	where := " where 1=1"

	if filter.ID != 0 {
		where += fmt.Sprintf(" AND id_cliente = %d", filter.ID)
	}
	if filter.Nome != "" {
		where += fmt.Sprintf(" AND nome ILIKE '%%%s%%'", filter.Nome)
	}

	if paginated {
		where += " ORDER BY " + filter.SortBy + " "
		if filter.Desc {
			where += " desc "
		} else {
			where += "asc"
		}
		where += ` limit ` + strconv.FormatInt(filter.Offset, 10)
		where += ` offset ` + strconv.FormatInt(filter.Page*filter.Offset, 10)
	}
	return where
}

// ListFuncionario returns a slice of public agents that earns more than the given pattern
func ClienteList(filter *ClienteFilter) ([]model.Cliente, error) {
	clienteList := []model.Cliente{}
	query := `SELECT id_cliente, nome FROM cliente ` + filter.makeFilter(true)
	rows, err := DBCon.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pf := model.Cliente{}

		err := rows.Scan(&pf.ID, &pf.Nome)
		if err != nil {
			return nil, err
		}
		clienteList = append(clienteList, pf)
	}
	return clienteList, nil
}
