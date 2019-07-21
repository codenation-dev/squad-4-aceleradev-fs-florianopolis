package adding

import (
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

type Service interface {
	AddUser(entity.User) error
	AddPublicFunc(pp ...entity.PublicFunc) error
	AddCustomer(cc ...entity.Customer) error

	ImportPublicFunc(mapFilter map[string][]string) error
	ImportCustomer() error

	// CreatePublicFuncTable(tableName string) error
	// CreateCustomerTable(tableName string) error
}

type Repository interface {
	CreateUser(entity.User) error
	CreateCustomer(cc ...entity.Customer) error
	CreatePublicFunc(pp ...entity.PublicFunc) error

	ImportPublicFunc(month, year string) error
	ImportCustomer() error

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

func (s *service) AddPublicFunc(pp ...entity.PublicFunc) error {
	return s.bR.CreatePublicFunc(pp...)
}

func (s *service) AddCustomer(cc ...entity.Customer) error {
	return s.bR.CreateCustomer(cc...)
}

func (s *service) ImportPublicFunc(mapFilter map[string][]string) error {
	month := "Maio"
	year := "2019"
	//TODO: implementar um esquema de escolher o mês usando queries para ver meses anteriores?
	// Ver como buscar sempre o latest
	return s.bR.ImportPublicFunc(month, year)
}

func (s *service) ImportCustomer() error {
	return s.bR.ImportCustomer()
}

// func (s *service) CreatePublicFuncTable(tableName string) error {
// 	return s.bR.CreatePublicFuncTable(tableName)
// }

// func (s *service) CreateCustomerTable(tableName string) error {
// 	return s.bR.CreateCustomerTable(tableName)
// }
