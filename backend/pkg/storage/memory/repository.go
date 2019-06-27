package memory

import (
	"codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"
	"errors"
)

// Storage keeps data in memory
type Storage struct {
	customers   []entity.Customer
	warnings    []entity.Warning
	users       []entity.User
	publicFuncs []entity.PublicFunc
}

func (m *Storage) AddCustomer(c entity.Customer) error {
	newC := entity.Customer{
		ID:          len(m.customers) + 1,
		Name:        c.Name,
		Wage:        c.Wage,
		IsPublic:    c.IsPublic,
		SentWarning: c.SentWarning,
	}
	m.customers = append(m.customers, newC)
	return nil
}

func (m *Storage) GetCustomerById(id int) (entity.Customer, error) {
	var customer entity.Customer

	for _, c := range m.customers {
		if c.ID == id {
			return c, nil
		}
	}
	return customer, errors.New("customer not found")
}

func (m *Storage) GetAllCustomers() ([]entity.Customer, error) {
	return m.customers, nil
}

// TODO: isso pode ser o mock do BD para os testes?
