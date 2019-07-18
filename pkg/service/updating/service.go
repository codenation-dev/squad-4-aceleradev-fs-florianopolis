package updating

import "github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"

type Service interface {
	ChangePassword(u entity.User) error
}

type Repository interface {
	UpdateUser(u entity.User) error
}

type service struct {
	bR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ChangePassword(u entity.User) error {
	return s.bR.UpdateUser(u)
}
