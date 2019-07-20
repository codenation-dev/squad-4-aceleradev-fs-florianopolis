package importing

import (
	"fmt"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

// FetchCustomerData imports the data from clientes.csv file
func FetchCustomerData(company string) ([]entity.Customer, error) {
	if company == "uati" {
		return fetchUatiData()
	}
	return nil, entity.ErrNotImplemented
}

func fetchUatiData() ([]entity.Customer, error) {
	var customers []entity.Customer
	path := "../cmd/data/downloaded/clientes.csv"
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
