// Package reading implements the rules to get data
package reading

import (
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/gorilla/schema"
)

// Service implements the business rules
type Service interface {
	GetUser(email string) (entity.User, error)
	GetAllUsers() ([]string, error)
	GetPublicFunc(mapFilter map[string][]string) (interface{}, error)
	GetCustomer(mapFilter map[string][]string) ([]entity.Customer, error)
	Query(q, offset, page string) (interface{}, error)
	StatsPublicFunc(mapFilter map[string][]string) (interface{}, error)
	DistPublicFunc(mapFilter map[string][]string) (map[string][]entity.PublicStats, error)

	// CompareCustomerPublicFunc(uf, month, year, company string) ([]entity.PublicFunc, error)
	// GetPublicFuncByWage(uf, year, month, wage string) ([]entity.PublicFunc, error)
}

// Repository implements the interface to deal with the storage
type Repository interface {
	ReadUser(email string) (entity.User, error)
	ReadAllUsers() ([]string, error)
	ReadPublicFunc(filter FuncFilter) (interface{}, error)
	ReadCustomer(filter CustFilter) ([]entity.Customer, error)
	Query(q, offset, page string) (interface{}, error)
	StatsPublicFunc(filter FuncFilter) (interface{}, error)
	DistPublicFunc(filter FuncFilter) (map[string][]entity.PublicStats, error)

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

func (s *service) Query(q, offset, page string) (interface{}, error) {
	//TODO: do jeito que está só serve no SQL, tenho que deixar mais genérica
	return s.bR.Query(q, offset, page)
}

func (s *service) GetUser(email string) (entity.User, error) {
	return s.bR.ReadUser(email)
}

func (s *service) GetAllUsers() ([]string, error) {
	return s.bR.ReadAllUsers()
}

type FuncFilter struct {
	Offset int64  `schema:"offset"`
	Page   int64  `schema:"page"`
	SortBy string `schema:"sortby"`
	Desc   bool   `schema:"desc"`

	//User specific filters
	ID         int64  `schema:"id"`
	Nome       string `schema:"nome"`
	Cargo      string `schema:"cargo"`
	Orgao      string `schema:"orgao"`
	Salario    int64  `schema:"salario"`
	Relevancia int64  `schema:"relevancia"`
	Customer   string `schema:"customer"` // yes - no - both
}

func validateFilter(mapFilter map[string][]string) (FuncFilter, error) {
	filter := FuncFilter{}
	if err := schema.NewDecoder().Decode(&filter, mapFilter); err != nil {
		return filter, err
	}

	if filter.Offset == 0 || filter.Offset > 50 {
		filter.Offset = 50
	}

	if filter.SortBy == "" {
		filter.SortBy = "short_name"
	}

	if filter.Customer != "yes" && filter.Customer != "no" {
		filter.Customer = ""
	}

	return filter, nil

}

func (s *service) GetPublicFunc(mapFilter map[string][]string) (interface{}, error) {
	filter, err := validateFilter(mapFilter)
	if err != nil {
		return nil, err
	}
	return s.bR.ReadPublicFunc(filter)
}

func (s *service) StatsPublicFunc(mapFilter map[string][]string) (interface{}, error) {
	filter, err := validateFilter(mapFilter)
	if err != nil {
		return nil, err
	}
	return s.bR.StatsPublicFunc(filter)
}

func (s *service) DistPublicFunc(mapFilter map[string][]string) (map[string][]entity.PublicStats, error) {
	filter, err := validateFilter(mapFilter)
	if err != nil {
		return nil, err
	}
	return s.bR.DistPublicFunc(filter)
}

type CustFilter struct {
	Offset int64  `schema:"offset"`
	Page   int64  `schema:"page"`
	SortBy string `schema:"sortby"`
	Desc   bool   `schema:"desc"`

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
	if filter.SortBy == "" {
		filter.SortBy = "name"
	}

	return s.bR.ReadCustomer(filter)
}
