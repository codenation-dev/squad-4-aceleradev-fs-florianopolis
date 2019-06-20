package adding

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
	"log"
)

// Service provides adding operations
type Service interface {
	AddCustomer(...entity.Customer)
}

// Repository provides access to customer repo
type Repository interface {
	AddCustomer(entity.Customer) error
}

type service struct { //TODO: não entendi ainda o porquê dessa struct
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddCustomer(c ...entity.Customer) {
	//TODO: some validation
	for _, customer := range c {
		err := s.bR.AddCustomer(customer)
		if err != nil {
			log.Fatal(err)
		}
	}
}
