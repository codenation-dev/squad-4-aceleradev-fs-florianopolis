package adding

import "codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"

// Service provides adding operations
type Service interface {
	AddCustomer(entity.Customer) error
	AddUser(entity.User) error
	AddWarning(entity.Warning) error
	AddPublicFunc(entity.PublicFunc) error
	// LoadPublicFuncFile() error
}

// Repository provides access to customer repo
type Repository interface {
	AddCustomer(entity.Customer) error
	AddUser(entity.User) error
	AddWarning(entity.Warning) error
	AddPublicFunc(entity.PublicFunc) error
	// LoadPublicFuncFile() error
}

type service struct { //TODO: não entendi ainda o porquê dessa struct
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddCustomer(customer entity.Customer) error {
	//TODO: some validation
	return s.bR.AddCustomer(customer)
	// for _, customer := range c {
	// 	err := s.bR.AddCustomer(customer)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}

func (s *service) AddUser(user entity.User) error {
	//TODO: some validation
	return s.bR.AddUser(user)
}

func (s *service) AddWarning(warning entity.Warning) error {
	return s.bR.AddWarning(warning)
}

func (s *service) AddPublicFunc(p entity.PublicFunc) error {
	return s.bR.AddPublicFunc(p)
}

// func (s *service) LoadPublicFuncFile() error {
// 	return s.bR.LoadPublicFuncFile()
// }
