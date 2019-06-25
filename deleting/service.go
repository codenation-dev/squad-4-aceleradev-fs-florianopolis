package deleting

// Service provides deleting methods
type Service interface {
	DeleteCustomerByID(id int) error
	DeleteUserByID(id int) error
	DeleteWarningByID(id int) error
}

// Repository provides access to DB
type Repository interface {
	DeleteCustomerByID(id int) error
	DeleteUserByID(id int) error
	DeleteWarningByID(id int) error
}

type service struct {
	bR Repository
}

// NewService creates a deleting service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteCustomerByID(id int) error {
	return s.bR.DeleteCustomerByID(id)

}

func (s *service) DeleteUserByID(id int) error {
	return s.bR.DeleteUserByID(id)
}

func (s *service) DeleteWarningByID(id int) error {
	return s.bR.DeleteWarningByID(id)
}
