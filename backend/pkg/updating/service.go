package updating

import "codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"

// Service provides updating methods
type Service interface {
	UpdateCustomer(customer entity.Customer) error
	UpdateUser(user entity.User) error
	UpdateWarning(warning entity.Warning) error
}

//Repository provides access to bd
type Repository interface {
	UpdateCustomer(c entity.Customer) error
	UpdateUser(user entity.User) error
	UpdateWarning(warning entity.Warning) error
}

type service struct {
	bR Repository
}

// NewService creates an updating service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateCustomer(customer entity.Customer) error {
	return s.bR.UpdateCustomer(customer)
}

func (s *service) UpdateUser(user entity.User) error {
	return s.bR.UpdateUser(user)
}

func (s *service) UpdateWarning(warning entity.Warning) error {
	return s.bR.UpdateWarning(warning)
}
