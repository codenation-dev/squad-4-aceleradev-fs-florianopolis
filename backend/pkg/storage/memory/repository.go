package memory

import (
	"errors"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
)

// Storage keeps data in memory
type Storage struct {
	customers   []model.Customer
	warnings    []model.Warning
	users       []model.User
	publicFuncs []model.PublicFunc
}

func (m *Storage) AddCustomer(c model.Cliente) error {
	newC := model.Customer{
		ID:          len(m.customers) + 1,
		Name:        c.Name,
		Wage:        c.Wage,
		IsPublic:    c.IsPublic,
		SentWarning: c.SentWarning,
	}
	m.customers = append(m.customers, newC)
	return nil
}

func (m *Storage) GetCustomerById(id int) (model.Customer, error) {
	var customer model.Customer

	for _, c := range m.customers {
		if c.ID == id {
			return c, nil
		}
	}
	return customer, errors.New("customer not found")
}

func (m *Storage) GetAllCustomers() ([]model.Customer, error) {
	return m.customers, nil
}

// TODO: isso pode ser o mock do BD para os testes?

// // Package memory implements the repository only in memory
// package memory

// import (
// 	"time"

// 	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/adding"
// 	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
// )

// // Storage implements the struct of the repo
// type Storage struct {
// 	customers []model.Customer
// }

// // CreateCustomer saves a new customer to repo
// func (m *Storage) CreateCustomer(c model.Customer) error {

// 	for _, savedCustomer := range m.customers {
// 		if savedCustomer.Name == c.Name {
// 			return adding.ErrDuplicate("cliente")
// 		}
// 	}
// 	c.ID = int(time.Now().Unix())
// 	m.customers = append(m.customers, c)
// 	return nil
// }
