package adding

import (
	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

type Service interface {
	AddUser(entity.User) error
	AddPublicFunc(tableName string, pp ...entity.PublicFunc) error
	AddCustomer(tableName string, cc ...entity.Customer) error

	// CreatePublicFuncTable(tableName string) error
	// CreateCustomerTable(tableName string) error
}

type Repository interface {
	CreateUser(entity.User) error
	CreateCustomer(tableName string, cc ...entity.Customer) error
	CreatePublicFunc(tableName string, pp ...entity.PublicFunc) error

	// CreatePublicFuncTable(tableName string) error
	// CreateCustomerTable(tableName string) error
}

type service struct {
	bR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddUser(u entity.User) error {
	u, err := entity.ValidateNewUser(u)
	if err != nil {
		return err
	}
	return s.bR.CreateUser(u)
}

func (s *service) AddPublicFunc(tableName string, pp ...entity.PublicFunc) error {
	return s.bR.CreatePublicFunc(tableName, pp...)
}

func (s *service) AddCustomer(tableName string, cc ...entity.Customer) error {
	return s.bR.CreateCustomer(tableName, cc...)
}

// func (s *service) CreatePublicFuncTable(tableName string) error {
// 	return s.bR.CreatePublicFuncTable(tableName)
// }

// func (s *service) CreateCustomerTable(tableName string) error {
// 	return s.bR.CreateCustomerTable(tableName)
// }
