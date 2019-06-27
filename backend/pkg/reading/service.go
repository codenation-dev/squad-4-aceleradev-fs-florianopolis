package reading

import "codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"

// Service provides reading operations
type Service interface {
	GetAllCustomers() ([]entity.Customer, error)
	GetAllUsers() ([]entity.User, error)
	GetAllWarnings() ([]entity.Warning, error)

	GetCustomerByID(id int) (entity.Customer, error)
	GetUserByID(id int) (entity.User, error)
	GetWarningByID(id int) (entity.Warning, error)
	GetPublicByID(id int) (entity.PublicFunc, error)

	GetCustomerByName(pattern string) ([]entity.Customer, error)
	GetUserByEmail(pattern string) ([]entity.User, error)
	GetWarningByCustomer(pattern string) ([]entity.Warning, error)
	GetWarningByUser(pattern string) ([]entity.Warning, error)

	GetPublicByWage(pattern float32) ([]entity.PublicFunc, error)
}

// Repository provides access to BD
type Repository interface {
	GetAllCustomers() ([]entity.Customer, error)
	GetAllUsers() ([]entity.User, error)
	GetAllWarnings() ([]entity.Warning, error)

	GetCustomerByID(id int) (entity.Customer, error)
	GetUserByID(id int) (entity.User, error)
	GetWarningByID(id int) (entity.Warning, error)
	GetPublicByID(id int) (entity.PublicFunc, error)

	GetCustomerByName(pattern string) ([]entity.Customer, error)
	GetUserByEmail(pattern string) ([]entity.User, error)
	GetWarningByCustomer(pattern string) ([]entity.Warning, error)
	GetWarningByUser(pattern string) ([]entity.Warning, error)

	GetPublicByWage(pattern float32) ([]entity.PublicFunc, error)
}

type service struct {
	bR Repository
}

// NewService creates a reading service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// All

// GetAllCustomers implements method
func (s *service) GetAllCustomers() ([]entity.Customer, error) {
	return s.bR.GetAllCustomers()
}

// GetAllUsers implements method
func (s *service) GetAllUsers() ([]entity.User, error) {
	return s.bR.GetAllUsers()
}

func (s *service) GetAllWarnings() ([]entity.Warning, error) {
	return s.bR.GetAllWarnings()
}

// ById

// GetCustomerByID implements method
func (s *service) GetCustomerByID(id int) (entity.Customer, error) {
	return s.bR.GetCustomerByID(id)
}

// GetUserByID implements method
func (s *service) GetUserByID(id int) (entity.User, error) {
	return s.bR.GetUserByID(id)
}

// GetWarningByID implements method
func (s *service) GetWarningByID(id int) (entity.Warning, error) {
	return s.bR.GetWarningByID(id)
}

func (s *service) GetPublicByID(id int) (entity.PublicFunc, error) {
	return s.bR.GetPublicByID(id)
}

// ByName

// GetCustomerByName implements method
func (s *service) GetCustomerByName(pattern string) ([]entity.Customer, error) {
	return s.bR.GetCustomerByName(pattern)
}

// GetUserByEmail implements method
func (s *service) GetUserByEmail(pattern string) ([]entity.User, error) {
	return s.bR.GetUserByEmail(pattern)
}

func (s *service) GetWarningByCustomer(pattern string) ([]entity.Warning, error) {
	return s.bR.GetWarningByCustomer(pattern)
}

func (s *service) GetWarningByUser(pattern string) ([]entity.Warning, error) {
	return s.bR.GetWarningByUser(pattern)
}

func (s *service) GetPublicByWage(pattern float32) ([]entity.PublicFunc, error) {
	return s.bR.GetPublicByWage(pattern)
}
