package database

import (
	"fmt"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
)

type FuncionarioFilter struct {
	Offset int64  `schema:"offset"`
	Page   int64  `schema:"page"`
	SortBy string `schema:"sortby"`
	Desc   bool   `schema:"asc"`

	//User specific filters
	ID      int64  `schema:"id"`
	Nome    string `schema:"nome"`
	Cargo   string `schema:"cargo"`
	Orgao   string `schema:"orgao"`
	Salario int64  `schema:"salario"`
}

func (filter FuncionarioFilter) makeFilter(paginated bool) string {
	where := " where 1 = 1"

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

// ListFuncionario returns a slice of public agents that earns more than the given pattern
func FuncionarioList(filter *FuncionarioFilter) ([]model.Funcionario, error) {
	funcionarioList := []model.Funcionario{}
	query := `SELECT nome, cargo, orgao, salario_mensal FROM funcionario ` + filter.makeFilter(true)
	fmt.Println(query)
	rows, err := DBCon.Query(query)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pf := model.Funcionario{}

		err := rows.Scan(&pf.Nome, &pf.Cargo, &pf.Orgao, &pf.SalarioMensal)
		if err != nil {
			return nil, err
		}
		funcionarioList = append(funcionarioList, pf)
	}
	fmt.Println(funcionarioList)
	return funcionarioList, nil
}

// GetPublicByID read a public_func from the DB, given the id
func FuncionarioGet(filter *FuncionarioFilter) (model.Funcionario, error) {
	query := `SELECT * FROM funcionario ` + filter.makeFilter(false)
	fmt.Println(query)
	funcionario := model.Funcionario{}
	err := DBCon.QueryRow(query).Scan(&funcionario.ID, &funcionario.MesReferencia, &funcionario.Nome, &funcionario.NomePesquisa,
		&funcionario.Cargo, &funcionario.Orgao, &funcionario.Estado, &funcionario.SalarioMensal, &funcionario.SalarioFerias,
		&funcionario.PagtoEventual, &funcionario.LicencaPremio, &funcionario.AbonoSalario, &funcionario.RedutorSalarial, &funcionario.TotalLiquido)

	if err != nil {
		return model.Funcionario{}, err
	}
	return funcionario, err
}
