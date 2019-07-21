// Package updating implementa as interfaces para modificar informações do banco de dados
package updating

import "github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"

// Service provides updating methods
type Service interface {
	UpdateCustomer(customer model.Customer) error
	UpdateUser(user model.User) error
	UpdateWarning(warning model.Warning) error
}

//Repository provides access to bd
type Repository interface {
	UpdateCustomer(c model.Customer) error
	UpdateUser(user model.User) error
	UpdateWarning(warning model.Warning) error
}

type service struct {
	bR Repository
}

// NewService creates an updating service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateCustomer(customer model.Customer) error {
	return s.bR.UpdateCustomer(customer)
}

func (s *service) UpdateUser(user model.User) error {
	return s.bR.UpdateUser(user)
}

func (s *service) UpdateWarning(warning model.Warning) error {
	return s.bR.UpdateWarning(warning)
}
