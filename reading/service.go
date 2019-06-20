package reading

// Service provides reading operations
type Service interface {
	GetAllCustomers() ([]Customer, error)
	GetCustomerByID(Customer) (Customer, error)
}

// Repository provides access to BD
type Repository interface {
	GetAllCustomers() ([]Customer, error)
	GetCustomerByID(Customer) (Customer, error)
}

type service struct {
	bR Repository
}

// NewService creates a reading service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAllCustomers() ([]Customer, error) {
	return s.bR.GetAllCustomers()
}

func (s *service) GetCustomerByID(c Customer) (Customer, error) {
	return s.bR.GetCustomerByID(c)
}
