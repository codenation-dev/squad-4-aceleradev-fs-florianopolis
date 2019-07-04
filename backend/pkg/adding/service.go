// Package adding implementa as interfaces para adicionar informações ao banco de dados
package adding

import (
	"fmt"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"
)

// ErrDuplicate is used when data already exists
func ErrDuplicate(data string) error {
	return fmt.Errorf("%s já existe", data)
}

// Service provides adding operations
type Service interface {
	AddCustomer(entity.Customer) error
	AddUser(entity.User) error
	AddWarning(entity.Warning) error
	AddPublicFunc(...entity.PublicFunc) error
	// LoadPublicFuncFile() error
}

// Repository provides access to customer repo
type Repository interface {
	AddCustomer(entity.Customer) error
	AddUser(entity.User) error
	AddWarning(entity.Warning) error
	AddPublicFunc(...entity.PublicFunc) error
	// LoadPublicFuncFile() error
}

type service struct {
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddCustomer(c entity.Customer) error {
	//TODO: some validation
	c.Name = entity.FormatString(c.Name)
	return s.bR.AddCustomer(c)
}

func (s *service) AddUser(user entity.User) error {
	//TODO: some validation
	return s.bR.AddUser(user)
}

func (s *service) AddWarning(warning entity.Warning) error {
	return s.bR.AddWarning(warning)
}

func (s *service) AddPublicFunc(pp ...entity.PublicFunc) error {
	return s.bR.AddPublicFunc(pp...)
}

// func (s *service) LoadPublicFuncFile() error {
// 	return s.bR.LoadPublicFuncFile()
// }
