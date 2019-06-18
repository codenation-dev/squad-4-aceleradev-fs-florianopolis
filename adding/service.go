package adding

import (
	"log"
)

// Service provides adding operations
type Service interface {
	AddCustomer(...Customer)
}

// Repository provides access to customer repo
type Repository interface {
	AddCustomer(Customer) error
}

type service struct { //TODO: não entendi ainda o porquê dessa struct
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddCustomer(c ...Customer) {
	//TODO: some validation
	for _, customer := range c {
		err := s.bR.AddCustomer(customer)
		if err != nil {
			log.Fatal(err)
		}
	}
}
