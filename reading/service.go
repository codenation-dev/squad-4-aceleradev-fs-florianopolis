package reading

import "codenation/squad-4-aceleradev-fs-florianopolis/entity"

// Service provides reading operations
type Service interface {
	GetAllCustomers() ([]entity.Customer, error)
	GetAllUsers() ([]entity.User, error)

	GetCustomerByID(id int) (entity.Customer, error)
	GetUserByID(id int) (entity.User, error)

	GetCustomerByName(pattern string) ([]entity.Customer, error)
	GetUserByEmail(pattern string) ([]entity.User, error)
}

// Repository provides access to BD
type Repository interface {
	GetAllCustomers() ([]entity.Customer, error)
	GetAllUsers() ([]entity.User, error)

	GetCustomerByID(id int) (entity.Customer, error)
	GetUserByID(id int) (entity.User, error)

	GetCustomerByName(pattern string) ([]entity.Customer, error)
	GetUserByEmail(pattern string) ([]entity.User, error)
}

type service struct {
	bR Repository
}

// NewService creates a reading service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetAllCustomers implements method
func (s *service) GetAllCustomers() ([]entity.Customer, error) {
	return s.bR.GetAllCustomers()
}

// GetAllUsers implements method
func (s *service) GetAllUsers() ([]entity.User, error) {
	return s.bR.GetAllUsers()
}

// GetCustomerByID implements method
func (s *service) GetCustomerByID(id int) (entity.Customer, error) {
	return s.bR.GetCustomerByID(id)
}

// GetUserByID implements method
func (s *service) GetUserByID(id int) (entity.User, error) {
	return s.bR.GetUserByID(id)
}

// GetCustomerByName implements method
func (s *service) GetCustomerByName(pattern string) ([]entity.Customer, error) {
	return s.bR.GetCustomerByName(pattern)
}

// GetUserByEmail implements method
func (s *service) GetUserByEmail(pattern string) ([]entity.User, error) {
	return s.bR.GetUserByEmail(pattern)
}
