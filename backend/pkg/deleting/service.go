// Package deleting implementa as interfaces para deletar informações do banco de dados
package deleting

// Service provides deleting methods
type Service interface {
	DeleteCustomerByID(id int) (int64, error)
	DeleteUserByID(id int) (int64, error)
	DeleteWarningByID(id int) (int64, error)
	DeletePublicByID(id int) (int64, error)
}

// Repository provides access to DB
type Repository interface {
	DeleteCustomerByID(id int) (int64, error)
	DeleteUserByID(id int) (int64, error)
	DeleteWarningByID(id int) (int64, error)
	DeletePublicByID(id int) (int64, error)
}

type service struct {
	bR Repository
}

// NewService creates a deleting service with all dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteCustomerByID(id int) (int64, error) {
	return s.bR.DeleteCustomerByID(id)

}

func (s *service) DeleteUserByID(id int) (int64, error) {
	return s.bR.DeleteUserByID(id)
}

func (s *service) DeleteWarningByID(id int) (int64, error) {
	return s.bR.DeleteWarningByID(id)
}

func (s *service) DeletePublicByID(id int) (int64, error) {
	return s.bR.DeletePublicByID(id)
}
