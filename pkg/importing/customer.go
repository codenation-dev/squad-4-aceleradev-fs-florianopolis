package importing

import (
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

func ImportCustomer() ([]entity.Customer, error) {
	var customers []entity.Customer
	path := entity.CacheFolder + "clientes.csv"
	job := func(row []string) bool {
		customers = append(customers, entity.Customer{
			Name: row[0],
		})
		return true
	}

	readCSV(path, job, ',', true)
	return customers, nil
}
