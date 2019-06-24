package updating

import "codenation/squad-4-aceleradev-fs-florianopolis/entity"

// Service provides updating methods
type Service interface {
	UpdateCustomer(customer entity.Customer) error
	UpdateUser(user entity.User) error
}

//Repository provides access to bd
type Repository interface {
	UpdateCustomer(c entity.Customer) error
	UpdateUser(user entity.User) error
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
	// for _, c := range customers {
	// 	err := s.bR.UpdateCustomer(c)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// return nil
}

func (s *service) UpdateUser(user entity.User) error {
	return s.bR.UpdateUser(user)
}
