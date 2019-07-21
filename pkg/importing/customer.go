package importing

import (
	"fmt"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

func importCustomerData(path string) ([]entity.Customer, error) {
	var customers []entity.Customer
	path = "../cmd/data/cache/clientes.csv"
	job := func(row []string) bool {
		customers = append(customers, entity.Customer{
			Name: row[0],
		})
		return true
	}

	readCSV(path, job, ',', true)
	fmt.Println(customers)
	return customers, nil
}
