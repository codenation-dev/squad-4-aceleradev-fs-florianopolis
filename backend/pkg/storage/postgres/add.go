package postgres

import (
	"fmt"
	"log"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
)

// AddCustomer inserts a new customer on the DB
func (s *Storage) AddCustomer(c model.Cliente) error {
	_, err := s.db.Exec(`INSERT INTO cliente (nome, nome_pesquisa)
						VALUES ($1, $2)`,
		&c.Nome, &c.NomePesquisa)
	fmt.Println(c, err)

	return err
}

// AddUser inserts a new user on the DB
func (s *Storage) AddUser(u model.User) error {
	bPass, err := model.Bcrypt(u.Pass)
	if err != nil {
		return err
	}
	u.Pass = string(bPass)
	_, err = s.db.Exec(`INSERT INTO users (email, pass)
						VALUES ($1, $2)`,
		&u.Email, &u.Pass)
	return err
}

// AddWarning inserts a new warning on the DB
func (s *Storage) AddWarning(w model.Warning) error {
	_, err := s.db.Exec(`INSERT INTO warnings (dt, msg, sent_to, from_customer)
						VALUES ($1, $2, $3, $4)`,
		&w.Dt, &w.Message, &w.SentTo, &w.FromCustomer)
	return err
}

// AddPublicFunc inserts a new public agent on the DB
func (s *Storage) AddPublicFunc(pp ...model.Funcionario) error {
	var insert = `INSERT INTO funcionario (mes_referencia, nome, nome_pesquisa, cargo, orgao, estado,
		salario_mensal, salario_ferias, pagto_eventual, licenca_premio, abono_salario, redutor_salarial, total_liquido) VALUES `
	var query = insert
	var vals = []interface{}{}
	batch := 5000
	i := 0
	for _, p := range pp {
		// we cannot use the '?' in the query because limitations of the driver
		// so we used the '$1, $2, $3...' notation
		inc := fmt.Sprintf("($%v, $%v, $%v,$%v, $%v, $%v,$%v, $%v, $%v,$%v, $%v, $%v, $%v), ",
			(1 + (i * 13)), (2 + (i * 13)), (3 + (i * 13)),
			(4 + (i * 13)), (5 + (i * 13)), (6 + (i * 13)),
			(7 + (i * 13)), (8 + (i * 13)), (9 + (i * 13)),
			(10 + (i * 13)), (11 + (i * 13)), (12 + (i * 13)),
			(13 + (i * 13)))
		i++
		query += inc
		vals = append(vals, p.MesReferencia, p.Nome, p.NomePesquisa, p.Cargo, p.Orgao, p.Estado,
			p.SalarioMensal, p.SalarioFerias, p.PagtoEventual, p.LicencaPremio, p.AbonoSalario, p.RedutorSalarial, p.TotalLiquido)

		if i%batch == 0 && i != 0 {
			q := query[0 : len(query)-2]
			_, err := s.db.Exec(q, vals...)
			if err != nil {
				log.Fatalf("error executing batch:", err)
			}
			// restart the vars to a new batch
			query = insert
			vals = []interface{}{}
			i = 0
		}
	}
	q := query[0 : len(query)-2]

	_, err := s.db.Exec(q, vals...)
	if err != nil {
		log.Fatalf("error executing the remaining batch", err)
	}

	return err
}
