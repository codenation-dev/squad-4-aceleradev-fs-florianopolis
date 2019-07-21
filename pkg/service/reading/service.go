// Package reading implements the rules to get data
package reading

import (
	"fmt"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/gorilla/schema"
)

// Service implements the business rules
type Service interface {
	GetUser(email string) (entity.User, error)
	GetPublicFunc(mapFilter map[string][]string) ([]entity.PublicFunc, error)
	GetCustomer(mapFilter map[string][]string) ([]entity.Customer, error)

	// CompareCustomerPublicFunc(uf, month, year, company string) ([]entity.PublicFunc, error)
	// GetPublicFuncByWage(uf, year, month, wage string) ([]entity.PublicFunc, error)
}

// Repository implements the interface to deal with the storage
type Repository interface {
	ReadUser(email string) (entity.User, error)
	ReadPublicFunc(filter FuncFilter) ([]entity.PublicFunc, error)
	ReadCustomer(filter CustFilter) ([]entity.Customer, error)

	// CompareCustomerPublicFunc(funcTableName, customerTableName string) ([]entity.PublicFunc, error)
	// ReadPublicFuncByWage(tableName, wage string) ([]entity.PublicFunc, error)
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

type FuncFilter struct {
	Offset int64  `schema:"offset"`
	Page   int64  `schema:"page"`
	SortBy string `schema:"sortby"`
	Desc   bool   `schema:"asc"`

	//User specific filters
	ID      int64  `schema:"id"`
	Nome    string `schema:"nome"`
	Cargo   string `schema:"cargo"`
	Orgao   string `schema:"orgao"`
	Salario int64  `schema:"salario"`
}

func (s *service) GetPublicFunc(mapFilter map[string][]string) ([]entity.PublicFunc, error) {
	filter := FuncFilter{}
	if err := schema.NewDecoder().Decode(&filter, mapFilter); err != nil {
		return nil, err
	}

	if filter.Offset == 0 || filter.Offset > 50 {
		filter.Offset = 50
	}

	if filter.SortBy == "" {
		filter.SortBy = "name"
	}

	return s.bR.ReadPublicFunc(filter)
}

type CustFilter struct {
	Offset int64  `schema:"offset"`
	Page   int64  `schema:"page"`
	SortBy string `schema:"sortby"`
	Desc   bool   `schema:"asc"`

	//User specific filters
	ID   int64  `schema:"id"`
	Name string `schema:"name"`
}

func (s *service) GetCustomer(mapFilter map[string][]string) ([]entity.Customer, error) {
	filter := CustFilter{}
	err := schema.NewDecoder().Decode(&filter, mapFilter)
	if err != nil {
		return nil, err
	}
	if filter.Offset == 0 || filter.Offset > 50 {
		filter.Offset = 50
	}
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.SortBy == "" {
		filter.SortBy = "name"
	}
	// Do something with filter
	fmt.Printf("%+v", filter)

	return s.bR.ReadCustomer(filter)
}

// func (s *service) CompareCustomerPublicFunc(uf, month, year, company string) ([]entity.PublicFunc, error) {
// 	funcTableName := makeTablename(uf, year, month)
// 	customerTableName := company
// 	return s.bR.CompareCustomerPublicFunc(funcTableName, customerTableName)
// }

// func (s *service) GetPublicFuncByWage(uf, year, month, wage string) ([]entity.PublicFunc, error) {
// 	tableName := makeTablename(uf, year, month)
// 	return s.bR.ReadPublicFuncByWage(tableName, wage)
// }

// func makeTablename(uf, year, month string) string {
// 	return fmt.Sprintf("public_func_%s_%s_%s", uf, year, month)
// }
