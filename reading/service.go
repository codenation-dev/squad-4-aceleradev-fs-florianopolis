package reading

import "codenation/squad-4-aceleradev-fs-florianopolis/entity"

// Service provides reading operations
type Service interface {
	GetAllCustomers() ([]entity.Customer, error)
	GetCustomerByID(id int) (entity.Customer, error)
	GetCustomerByName(pattern string) ([]entity.Customer, error)
}

// Repository provides access to BD
type Repository interface {
	GetAllCustomers() ([]entity.Customer, error)
	GetCustomerByID(id int) (entity.Customer, error)
	GetCustomerByName(pattern string) ([]entity.Customer, error)
}

type service struct {
	bR Repository
}
 
// NewService creates a reading service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAllCustomers() ([]entity.Customer, error) {
	return s.bR.GetAllCustomers()
}

func (s *service) GetCustomerByID(id int) (entity.Customer, error) {
	return s.bR.GetCustomerByID(id)
}

func (s *service) GetCustomerByName(pattern string) ([]entity.Customer, error) {
	return s.bR.GetCustomerByName(pattern)
}
