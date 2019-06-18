package reading

// Service provides reading operations
type Service interface {
	GetAllCustomers() []Customer
}

// Repository provides access to BD
type Repository interface {
	GetAllCustomers() ([]Customer, error)
}

type service struct {
	bR Repository
}

// NewService creates a reading service with all dependencies
func NewService(r Repository) Service {
	return &service(r) //TODO: por que o adding aceita e o reading não???
}

func (s *service) GetAllCustomers() []Customer {
	c, err := s.bR.GetAllCustomers()
	if err != nil {
		return nil
	}
	return c
}
