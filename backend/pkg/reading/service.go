package reading

import "github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"

// Service provides reading operations
type Service interface {
	GetAllCustomers() ([]model.Customer, error)
	GetAllUsers() ([]model.User, error)
	GetAllWarnings() ([]model.Warning, error)

	GetCustomerByID(id int) (model.Customer, error)
	GetUserByID(id int) (model.User, error)
	GetWarningByID(id int) (model.Warning, error)
	GetPublicByID(id int) (model.PublicFunc, error)

	GetCustomerByName(pattern string) ([]model.Customer, error)
	GetUserByEmail(pattern string) (model.User, error)
	GetWarningByCustomer(pattern string) ([]model.Warning, error)
	GetWarningByUser(pattern string) ([]model.Warning, error)

	GetPublicByWage(pattern float32) ([]model.PublicFunc, error)
}

// Repository provides access to BD
type Repository interface {
	GetAllCustomers() ([]model.Customer, error)
	GetAllUsers() ([]model.User, error)
	GetAllWarnings() ([]model.Warning, error)

	GetCustomerByID(id int) (model.Customer, error)
	GetUserByID(id int) (model.User, error)
	GetWarningByID(id int) (model.Warning, error)
	GetPublicByID(id int) (model.PublicFunc, error)

	GetCustomerByName(pattern string) ([]model.Customer, error)
	GetUserByEmail(pattern string) (model.User, error)
	GetWarningByCustomer(pattern string) ([]model.Warning, error)
	GetWarningByUser(pattern string) ([]model.Warning, error)

	GetPublicByWage(pattern float32) ([]model.PublicFunc, error)
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
func (s *service) GetAllCustomers() ([]model.Customer, error) {
	return s.bR.GetAllCustomers()
}

// GetAllUsers implements method
func (s *service) GetAllUsers() ([]model.User, error) {
	return s.bR.GetAllUsers()
}

func (s *service) GetAllWarnings() ([]model.Warning, error) {
	return s.bR.GetAllWarnings()
}

// ById

// GetCustomerByID implements method
func (s *service) GetCustomerByID(id int) (model.Customer, error) {
	return s.bR.GetCustomerByID(id)
}

// GetUserByID implements method
func (s *service) GetUserByID(id int) (model.User, error) {
	return s.bR.GetUserByID(id)
}

// GetWarningByID implements method
func (s *service) GetWarningByID(id int) (model.Warning, error) {
	return s.bR.GetWarningByID(id)
}

func (s *service) GetPublicByID(id int) (model.PublicFunc, error) {
	return s.bR.GetPublicByID(id)
}

// ByName

// GetCustomerByName implements method
func (s *service) GetCustomerByName(pattern string) ([]model.Customer, error) {
	return s.bR.GetCustomerByName(pattern)
}

// GetUserByEmail implements method
func (s *service) GetUserByEmail(pattern string) (model.User, error) {
	return s.bR.GetUserByEmail(pattern)
}

func (s *service) GetWarningByCustomer(pattern string) ([]model.Warning, error) {
	return s.bR.GetWarningByCustomer(pattern)
}

func (s *service) GetWarningByUser(pattern string) ([]model.Warning, error) {
	return s.bR.GetWarningByUser(pattern)
}

func (s *service) GetPublicByWage(pattern float32) ([]model.PublicFunc, error) {
	return s.bR.GetPublicByWage(pattern)
}
