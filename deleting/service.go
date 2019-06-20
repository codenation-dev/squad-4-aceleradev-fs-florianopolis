package deleting

// Service provides deleting methods
type Service interface {
	DeleteCustomerByID(ids ...int) error
}

// Repository provides access to DB
type Repository interface {
	DeleteCustomerByID(id int) error
}

type service struct {
	bR Repository
}

// NewService creates a deleting service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteCustomerByID(ids ...int) error {
	for _, id := range ids {
		err := s.bR.DeleteCustomerByID(id)
		//TODO: tratar erro aqui, na regra de negócio(?)
		if err != nil {
			return err
		}
	}
	return nil
}
