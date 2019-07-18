// Package reading implements the rules to get data
package reading

import (
	"fmt"

	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

// Service implements the business rules
type Service interface {
	GetUser(email string) (entity.User, error)
	GetAllPublicFunc(uf, year, month string) ([]entity.PublicFunc, error)
	GetAllCustomers(tableName string) ([]entity.Customer, error)

	CompareCustomerPublicFunc(uf, month, year, company string) ([]entity.PublicFunc, error)
}

// Repository implements the interface to deal with the storage
type Repository interface {
	ReadUser(email string) (entity.User, error)
	ReadAllPublicFunc(uf, year, month string) ([]entity.PublicFunc, error)
	ReadAllCustomers(tableName string) ([]entity.Customer, error)

	CompareCustomerPublicFunc(funcTableName, customerTableName string) ([]entity.PublicFunc, error)
}

type service struct {
	bR Repository
}

// NewService starts a new service with all the dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetUser(email string) (entity.User, error) {
	return s.bR.ReadUser(email)
}

func (s *service) GetAllPublicFunc(uf, year, month string) ([]entity.PublicFunc, error) {
	return s.bR.ReadAllPublicFunc(uf, year, month)
}

func (s *service) GetAllCustomers(company string) ([]entity.Customer, error) {
	tableName := company
	return s.bR.ReadAllCustomers(tableName)
}

func (s *service) CompareCustomerPublicFunc(uf, month, year, company string) ([]entity.PublicFunc, error) {
	funcTableName := fmt.Sprintf("public_func_%s_%s_%s", uf, year, month)
	customerTableName := company
	return s.bR.CompareCustomerPublicFunc(funcTableName, customerTableName)
}
