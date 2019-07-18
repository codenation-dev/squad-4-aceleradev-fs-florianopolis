package deleting

type Service interface {
	DeleteUser(email string) error
}

type Repository interface {
	DeleteUser(email string) error
}

type service struct {
	bR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteUser(email string) error {
	return s.bR.DeleteUser(email)
}
 