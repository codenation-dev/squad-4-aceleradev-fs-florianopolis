package updating

import "codenation/squad-4-aceleradev-fs-florianopolis/entity"

// Service provides updating methods
type Service interface {
	UpdateCustomer(customers ...entity.Customer) error
}

//Repository provides access to bd
type Repository interface {
	UpdateCustomer(c entity.Customer) error
}

type service struct {
	bR Repository
}

// NewService creates an updating service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateCustomer(customers ...entity.Customer) error {
	for _, c := range customers {
		err := s.bR.UpdateCustomer(c)
		if err != nil {
			return err
		}
	}
	return nil
}
